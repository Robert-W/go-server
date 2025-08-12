package sample

import (
	"context"
	"time"

	"github.com/robert-w/go-server/internal/constants"
	"github.com/robert-w/go-server/internal/monitoring"
	v1 "github.com/robert-w/go-server/internal/routes/v1"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	semconv "go.opentelemetry.io/otel/semconv/v1.34.0"
)

type sampleService struct {}

func (s *sampleService) listAllSamples(ctx context.Context) (*[]sample, *v1.Error) {
	span := monitoring.CreateDBSpan(ctx, "listAllSamples")
	defer span.End()

	samples := []sample{
		{
			Id:        "111",
			Value:     "First Sample",
			Timestamp: time.Now(),
		},
		{
			Id:        "222",
			Value:     "Second Sample",
			Timestamp: time.Now(),
		},
		{
			Id:        "333",
			Value:     "Third Sample",
			Timestamp: time.Now(),
		},
	}

	span.SetStatus(codes.Ok, "Ok")
	span.SetAttributes(
		semconv.DBResponseReturnedRows(len(samples)),
	)

	return &samples, nil
}

func (s *sampleService) createSamples(ctx context.Context) (*[]sample, *v1.Error) {
	span := monitoring.CreateDBSpan(ctx, "createSamples")
	defer span.End()

	samples := []sample{
		{
			Id:        "111",
			Value:     "New Sample",
			Timestamp: time.Now(),
		},
	}

	span.SetStatus(codes.Ok, "Ok")
	span.SetAttributes(
		semconv.DBResponseReturnedRows(len(samples)),
	)

	return &samples, nil
}

func (s *sampleService) getSampleById(ctx context.Context) (*sample, *v1.Error) {
	span := monitoring.CreateDBSpan(ctx, "getSampleById")
	defer span.End()

	sample := sample{
		Id:        "123",
		Value:     "Sample Read",
		Timestamp: time.Now(),
	}

	span.SetStatus(codes.Ok, "Ok")
	span.SetAttributes(
		semconv.DBResponseReturnedRows(1),
	)

	return &sample, nil
}

func (s *sampleService) updateSampleById(ctx context.Context) (*sample, *v1.Error) {
	span := monitoring.CreateDBSpan(ctx, "updateSampleById")
	defer span.End()

	sample := sample{
		Id:        "321",
		Value:     "Sample Update",
		Timestamp: time.Now(),
	}

	span.SetStatus(codes.Ok, "Ok")
	span.SetAttributes(
		attribute.Int(constants.DB_AFFECTED_ROWS, 1),
	)

	return &sample, nil
}

func (s *sampleService) deleteSampleById(ctx context.Context) (*sample, *v1.Error) {
	span := monitoring.CreateDBSpan(ctx, "deleteSampleById")
	defer span.End()

	sample := sample{
		Id:        "321",
	}

	span.SetStatus(codes.Ok, "Ok")
	span.SetAttributes(
		attribute.Int(constants.DB_AFFECTED_ROWS, 0),
	)

	return &sample, nil
}


