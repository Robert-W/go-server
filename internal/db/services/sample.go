package services

import (
	"context"
	"time"

	"github.com/robert-w/go-server/internal/constants"
	"github.com/robert-w/go-server/internal/monitoring"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	semconv "go.opentelemetry.io/otel/semconv/v1.34.0"
)

type Sample struct {
	Id        string    `json:"id"`
	Value     string    `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}

type SampleService struct {
	// TODO: add db pool here
}

func (s *SampleService) ListAllSamples(ctx context.Context) (*[]Sample, error) {
	span := monitoring.CreateDBSpan(ctx, "ListAllSamples")
	defer span.End()

	samples := []Sample{
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

func (s *SampleService) CreateSamples(ctx context.Context) (*[]Sample, error) {
	span := monitoring.CreateDBSpan(ctx, "CreateSamples")
	defer span.End()

	samples := []Sample{
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

func (s *SampleService) GetSampleById(ctx context.Context) (*Sample, error) {
	span := monitoring.CreateDBSpan(ctx, "GetSampleById")
	defer span.End()

	sample := Sample{
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

func (s *SampleService) UpdateSampleById(ctx context.Context) (*Sample, error) {
	span := monitoring.CreateDBSpan(ctx, "UpdateSampleById")
	defer span.End()

	sample := Sample{
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

func (s *SampleService) DeleteSampleById(ctx context.Context) error {
	span := monitoring.CreateDBSpan(ctx, "DeleteSampleById")
	defer span.End()

	span.SetStatus(codes.Ok, "Ok")
	span.SetAttributes(
		attribute.Int(constants.DB_AFFECTED_ROWS, 0),
	)

	return nil
}

