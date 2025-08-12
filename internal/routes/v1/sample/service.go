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

func (s *sampleService) ListAllSamples(ctx context.Context) (*[]sample, *v1.Error) {
	span := monitoring.CreateDBSpan(ctx, "ListAllSamples")
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

func (s *sampleService) CreateSamples(ctx context.Context) (*[]sample, *v1.Error) {
	span := monitoring.CreateDBSpan(ctx, "CreateSamples")
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

func (s *sampleService) GetSampleById(ctx context.Context) (*sample, *v1.Error) {
	span := monitoring.CreateDBSpan(ctx, "GetSampleById")
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

func (s *sampleService) UpdateSampleById(ctx context.Context) (*sample, *v1.Error) {
	span := monitoring.CreateDBSpan(ctx, "UpdateSampleById")
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

func (s *sampleService) DeleteSampleById(ctx context.Context) (*sample, *v1.Error) {
	span := monitoring.CreateDBSpan(ctx, "DeleteSampleById")
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


