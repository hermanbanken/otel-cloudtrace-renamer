module github.com/hermanbanken/otel-cloudtrace-renamer/example

replace github.com/hermanbanken/otel-cloudtrace-renamer => ../

go 1.16

require (
	github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/trace v1.0.0-RC2
	github.com/hermanbanken/otel-cloudtrace-renamer v0.0.0-00010101000000-000000000000
	github.com/jnovack/flag v1.16.0
	github.com/pkg/errors v0.9.1
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.22.0 // indirect
	go.opentelemetry.io/otel v1.0.0-RC2
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.0.0-RC2
	go.opentelemetry.io/otel/sdk v1.0.0-RC2
)
