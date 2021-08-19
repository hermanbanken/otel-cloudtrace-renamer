package main

import (
	"context"
	"fmt"
	"log"

	texporter "github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/trace"
	renamer "github.com/hermanbanken/otel-cloudtrace-renamer"
	"github.com/jnovack/flag"
	"github.com/pkg/errors"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

const GoogleCloudServiceKey = "g.co/gae/app/module"
const serviceName = "example-server"

// TelemetryExporter defines which export to send OpenTelemetry to
var TelemetryExporter = flag.String("telemetry-exporter", "google", "Set a telemetry exporter. Options are: stdout, google (future: jaeger?)")

// StartTelemetry provisions all telemetry settings so that other code can just create/use trace spans
func StartTelemetry(ctx context.Context, serviceName string) (err error) {
	// Set the network propagator format to W3C (https://www.w3.org/TR/trace-context/)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.Baggage{}, propagation.TraceContext{}))

	// Resource refers to this process unit (pod) and sets some default telemetry labels.
	// It supports setting extra tags using environment variable $OTEL_RESOURCE_ATTRIBUTES=key=value,foo=bar
	res, err := resource.New(context.Background(), resource.WithAttributes(
		attribute.String(GoogleCloudServiceKey, serviceName), // Cloud Trace displays this in the UI as "Service"; there is no alternative key yet
	))
	if err != nil {
		return errors.Wrap(err, "Failed to create telemetry resource")
	}

	// Exporter (traces)
	var exporterTraces sdktrace.SpanExporter
	switch *TelemetryExporter {
	default:
		return fmt.Errorf("invalid choice for -telemetry-exporter %q", *TelemetryExporter)
	case "none":
		exporterTraces = &noopExporter{}
	case "stdout":
		if exporterTraces, err = stdouttrace.New(stdouttrace.WithPrettyPrint()); err != nil {
			return errors.Wrap(err, "Failed to create stdout telemetry exporter")
		}
	case "google":
		if exporterTraces, err = texporter.New(); err != nil {
			return errors.Wrap(err, "Failed to create google telemetry exporter")
		}
		exporterTraces = renamer.CloudTraceAttributeRenamer{exporterTraces}
	}
	log.Printf("Sending telemetry to %q", *TelemetryExporter)

	// Trace provider
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithSpanProcessor(sdktrace.NewBatchSpanProcessor(exporterTraces)),
		sdktrace.WithResource(res),
	)
	otel.SetTracerProvider(tp)

	go func() {
		<-ctx.Done()
		// Were not handling these errors, because were already shutting down
		_ = tp.Shutdown(ctx)
		_ = exporterTraces.Shutdown(ctx)
	}()
	return nil
}

type noopExporter struct{}

var _ sdktrace.SpanExporter = &noopExporter{}

func (*noopExporter) ExportSpans(ctx context.Context, ss []sdktrace.ReadOnlySpan) error {
	return nil
}
func (*noopExporter) Shutdown(ctx context.Context) error {
	return nil
}
