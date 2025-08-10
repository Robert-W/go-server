package monitoring

import (
	"context"
	"os"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	oteltrace "go.opentelemetry.io/otel/trace"
)

func NewTracer(ctx context.Context) (oteltrace.Tracer, *sdktrace.TracerProvider, error) {
	tracer := otel.Tracer("api-server")

	exporter, err := otlptracehttp.New(
		ctx,
		otlptracehttp.WithCompression(otlptracehttp.GzipCompression),
		otlptracehttp.WithEndpoint("0.0.0.0:4318"),
		otlptracehttp.WithInsecure(),
	)
	if err != nil {
		return nil, nil, err
	}

	traceProvider := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(resource.NewSchemaless(
			attribute.String("service.name", "api-server"),
			attribute.Int("process.pid", os.Getpid()),
		)),
	)
	otel.SetTracerProvider(traceProvider)

	textMapPropagator := propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	)
	otel.SetTextMapPropagator(textMapPropagator)

	return tracer, traceProvider, nil
}
