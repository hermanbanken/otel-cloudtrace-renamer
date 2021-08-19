package renamer

import (
	"context"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

// CloudTraceAttributeRenamer converts from semconv to https://cloud.google.com/trace/docs/trace-labels format
type CloudTraceAttributeRenamer struct {
	sdktrace.SpanExporter
}

var _ sdktrace.SpanExporter = CloudTraceAttributeRenamer{}

func (t CloudTraceAttributeRenamer) ExportSpans(ctx context.Context, spans []sdktrace.ReadOnlySpan) error {
	for _, s := range spans {
		attrs := s.Attributes()
		for i, attr := range attrs {
			switch attr.Key {
			case semconv.HTTPMethodKey:
				attr.Key = "/http/method"
			case semconv.HTTPHostKey:
				attr.Key = "/http/host"
			case semconv.HTTPTargetKey:
				attr.Key = "/http/path"
			case semconv.HTTPStatusCodeKey:
				attr.Key = "/http/status_code"
			case semconv.HTTPResponseContentLengthKey, otelhttp.WroteBytesKey:
				attr.Key = "/http/response/size"
			case semconv.HTTPRequestContentLengthKey, otelhttp.ReadBytesKey:
				attr.Key = "/http/request/size"
			case semconv.HTTPRouteKey:
				attr.Key = "/http/route"
			case semconv.HTTPUserAgentKey:
				attr.Key = "/http/user_agent"

			case semconv.ExceptionMessageKey:
				attr.Key = "/error/message"
			case semconv.ExceptionEventName:
				attr.Key = "/error/name"
			case semconv.ExceptionStacktraceKey:
				attr.Key = "/stacktrace"
			}
			attrs[i] = attr
		}
	}
	return t.SpanExporter.ExportSpans(ctx, spans)
}
func (t CloudTraceAttributeRenamer) Shutdown(ctx context.Context) error {
	return t.SpanExporter.Shutdown(ctx)
}
