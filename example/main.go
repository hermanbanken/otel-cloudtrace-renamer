package main

import (
	"context"
	"net/http"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

const serviceName = "example-server"

func main() {
	StartTelemetry(context.Background(), serviceName)
	http.Handle("/", http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("Hello World"))
	}))
	http.ListenAndServe(":8080", otelhttp.NewHandler(http.DefaultServeMux, "server", otelhttp.WithPublicEndpoint()))
}
