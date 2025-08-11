package models

import (
	"context"
	"time"

	"github.com/robert-w/go-server/internal/constants"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
)

type Sample struct {
	Id        string    `json:"id"`
	Value     string    `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}

func ListAllSamples(ctx context.Context) (*[]Sample, error) {
	tracer := otel.Tracer(constants.SERVICE_NAME)
	_, span := tracer.Start(ctx, "ListAllSamples")
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

	return &samples, nil
}

func CreateSamples(ctx context.Context) (*[]Sample, error) {
	tracer := otel.Tracer(constants.SERVICE_NAME)
	_, span := tracer.Start(ctx, "CreateSamples")
	defer span.End()

	samples := []Sample{
		{
			Id:        "111",
			Value:     "New Sample",
			Timestamp: time.Now(),
		},
	}

	span.SetStatus(codes.Ok, "Ok")

	return &samples, nil
}

func GetSampleById(ctx context.Context) (*Sample, error) {
	tracer := otel.Tracer(constants.SERVICE_NAME)
	_, span := tracer.Start(ctx, "GetSampleById")
	defer span.End()

	sample := Sample{
		Id:        "123",
		Value:     "Sample Read",
		Timestamp: time.Now(),
	}

	span.SetStatus(codes.Ok, "Ok")

	return &sample, nil
}

func UpdateSampleById(ctx context.Context) (*Sample, error) {
	tracer := otel.Tracer(constants.SERVICE_NAME)
	_, span := tracer.Start(ctx, "listSamplesQuery")
	defer span.End()

	sample := Sample{
		Id:        "321",
		Value:     "Sample Update",
		Timestamp: time.Now(),
	}

	span.SetStatus(codes.Ok, "Ok")

	return &sample, nil
}

func DeleteSampleById(ctx context.Context) error {
	tracer := otel.Tracer(constants.SERVICE_NAME)
	_, span := tracer.Start(ctx, "DeleteSampleById")
	defer span.End()

	span.SetStatus(codes.Ok, "Ok")

	return nil
}
