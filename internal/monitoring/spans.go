package monitoring

import (
	"context"
	"runtime"

	"github.com/robert-w/go-server/internal/constants"
	"go.opentelemetry.io/otel"
	semconv "go.opentelemetry.io/otel/semconv/v1.34.0"
	"go.opentelemetry.io/otel/trace"
)

func CreateDBSpan(ctx context.Context, name string) trace.Span {
	tracer := otel.Tracer(constants.SERVICE_NAME)
	pointer, file, line, _ := runtime.Caller(1)

	// Samplers apparently only have access to attributes provided at the time of
	// creation. You can update the values later, but if you need to sample on an
	// attribute, add it here
	options := []trace.SpanStartOption{
		trace.WithSpanKind(trace.SpanKindInternal),
		trace.WithAttributes(
			semconv.CodeFilePath(file),
			semconv.CodeLineNumber(line),
			semconv.CodeFunctionName(runtime.FuncForPC(pointer).Name()),
			semconv.DBCollectionName(""),
			semconv.DBSystemNamePostgreSQL,
			semconv.DBOperationName(""),
		),
	}

	_, span := tracer.Start(ctx, name, options...)

	return span
}

func CreateSpan(ctx context.Context, name string) (context.Context, trace.Span) {
	tracer := otel.Tracer(constants.SERVICE_NAME)
	pointer, file, line, _ := runtime.Caller(1)

	// Samplers apparently only have access to attributes provided at the time of
	// creation. You can update the values later, but if you need to sample on an
	// attribute, add it here
	options := []trace.SpanStartOption{
		trace.WithSpanKind(trace.SpanKindInternal),
		trace.WithAttributes(
			semconv.CodeFilePath(file),
			semconv.CodeLineNumber(line),
			semconv.CodeFunctionName(runtime.FuncForPC(pointer).Name()),
		),
	}

	return tracer.Start(ctx, name, options...)
}
