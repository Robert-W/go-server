package monitoring

import (
	"context"
	"os"
	"runtime"

	"github.com/robert-w/go-server/internal/constants"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	semconv "go.opentelemetry.io/otel/semconv/v1.34.0"
	"go.opentelemetry.io/otel/trace"
)

func getCommmoAttributes() []attribute.KeyValue {
	// Set this value to 2 because this function should only be called by one of
	// the CreateSpan functions below. Caller two will identify the function that
	// called CreateSpan, which is what we are going for here
	pointer, file, line, _ := runtime.Caller(2)

	return []attribute.KeyValue{
		semconv.CodeFilePath(file),
		semconv.CodeLineNumber(line),
		semconv.CodeFunctionName(runtime.FuncForPC(pointer).Name()),
		semconv.ProcessRuntimeVersion(runtime.Version()),
		semconv.ServiceName(constants.SERVICE_NAME),
		semconv.ServiceVersion(os.Getenv("GIT_SHA")),
	}
}

func CreateDBSpan(ctx context.Context, name string) trace.Span {
	tracer := otel.Tracer(constants.SERVICE_NAME)

	// Samplers apparently only have access to attributes provided at the time of
	// creation. You can update the values later, but if you need to sample on an
	// attribute, add it here
	options := []trace.SpanStartOption{
		trace.WithSpanKind(trace.SpanKindInternal),
		trace.WithAttributes(getCommmoAttributes()...),
		trace.WithAttributes(
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

	// Samplers apparently only have access to attributes provided at the time of
	// creation. You can update the values later, but if you need to sample on an
	// attribute, add it here
	options := []trace.SpanStartOption{
		trace.WithSpanKind(trace.SpanKindInternal),
		trace.WithAttributes(getCommmoAttributes()...),
	}

	return tracer.Start(ctx, name, options...)
}
