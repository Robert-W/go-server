package sample

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

func listSamplesQuery(ctx context.Context) (*[]Sample, error) {
	tracer := otel.Tracer(constants.SERVICE_NAME)
	_, span := tracer.Start(ctx, "listSamplesQuery")
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

func createSamplesQuery(ctx context.Context) (*[]Sample, error) {
	tracer := otel.Tracer(constants.SERVICE_NAME)
	_, span := tracer.Start(ctx, "listSamplesQuery")
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

func readSampleQuery(ctx context.Context) (*Sample, error) {
	tracer := otel.Tracer(constants.SERVICE_NAME)
	_, span := tracer.Start(ctx, "listSamplesQuery")
	defer span.End()

	sample := Sample{
		Id:        "123",
		Value:     "Sample Read",
		Timestamp: time.Now(),
	}

	span.SetStatus(codes.Ok, "Ok")

	return &sample, nil
}

func updateSampleQuery(ctx context.Context) (*Sample, error) {
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

func deleteSampleQuery(ctx context.Context) error {
	tracer := otel.Tracer(constants.SERVICE_NAME)
	_, span := tracer.Start(ctx, "listSamplesQuery")
	defer span.End()

	span.SetStatus(codes.Ok, "Ok")

	return nil
}
