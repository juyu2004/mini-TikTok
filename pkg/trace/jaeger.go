package trace

import (
	"context"
	"os"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"go.opentelemetry.io/otel/sdk/trace"
)

func Init(service string) func(context.Context) error {
	endpoint := os.Getenv("JAEGER_ENDPOINT")
	exporter, _ := otlptracehttp.New(context.Background(), otlptracehttp.WithEndpointURL(endpoint))
	tp := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithResource(resource.NewWithAttributes(semconv.SchemaURL, semconv.ServiceName(service))),
	)
	otel.SetTracerProvider(tp)
	return tp.Shutdown
}
