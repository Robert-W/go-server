package monitoring

import (
	"context"
	"os"

	"github.com/robert-w/go-server/internal/constants"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.34.0"
)

func NewTraceProvider(ctx context.Context) (*sdktrace.TracerProvider, error) {
	exporter, err := otlptracehttp.New(
		ctx,
		otlptracehttp.WithCompression(otlptracehttp.GzipCompression),
		otlptracehttp.WithEndpoint(os.Getenv("OTEL_COLLECTOR_URL")),
		otlptracehttp.WithInsecure(),
	)
	if err != nil {
		return nil, err
	}

	traceProvider := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(resource.NewSchemaless(
			semconv.ServiceName(constants.SERVICE_NAME),
			semconv.ServiceVersion(os.Getenv("GIT_SHA")),
			attribute.Int("process.pid", os.Getpid()),
		)),
	)
	otel.SetTracerProvider(traceProvider)

	textMapPropagator := propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	)
	otel.SetTextMapPropagator(textMapPropagator)

	return traceProvider, nil
}
