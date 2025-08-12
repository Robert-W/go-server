package monitoring

import (
	"context"

	"github.com/robert-w/go-server/internal/constants"
	"go.opentelemetry.io/otel"
	semconv "go.opentelemetry.io/otel/semconv/v1.34.0"
	"go.opentelemetry.io/otel/trace"
)

func CreateDBSpan(ctx context.Context, name string) trace.Span {
	tracer := otel.Tracer(constants.SERVICE_NAME)

	// Samplers apparently only have access to attributes provided at the time of
	// creation. You can update the values later, but if you need to sample on an
	// attribute, add it here
	options := []trace.SpanStartOption{
		trace.WithSpanKind(trace.SpanKindInternal),
		trace.WithAttributes(
			semconv.DBCollectionName(""),
			semconv.DBSystemNamePostgreSQL,
			semconv.DBOperationName(""),
		),
	}

	_, span := tracer.Start(ctx, name, options...)

	return span
}
