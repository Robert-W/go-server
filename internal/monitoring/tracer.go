package monitoring

import (
	"os"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	stdout "go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	oteltrace "go.opentelemetry.io/otel/trace"
)

func NewTracer() (oteltrace.Tracer, *sdktrace.TracerProvider, error) {
	tracer := otel.Tracer("api-server")

	exporter, err := stdout.New(stdout.WithPrettyPrint())
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
