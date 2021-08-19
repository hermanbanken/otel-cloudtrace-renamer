package main

import (
	"context"
	"net/http"
)

func main() {
	StartTelemetry(context.Background(), serviceName)
	http.Handle("/", http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("Hello World"))
	}))

	http.ListenAndServe(":8080", http.DefaultServeMux)
}
