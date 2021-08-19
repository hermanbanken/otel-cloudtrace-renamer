module github.com/hermanbanken/otel-cloudtrace-renamer

go 1.16

require (
	cloud.google.com/go v0.93.3 // indirect
	cloud.google.com/go/trace v0.1.0 // indirect
	github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/trace v1.0.0-RC2
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/jnovack/flag v1.16.0
	github.com/pkg/errors v0.9.1
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.22.0
	go.opentelemetry.io/otel v1.0.0-RC2
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.0.0-RC2
	go.opentelemetry.io/otel/sdk v1.0.0-RC2
)
