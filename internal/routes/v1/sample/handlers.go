package sample

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	oteltrace "go.opentelemetry.io/otel/trace"
)

type Sample struct {
	Id        string    `json:"id"`
	Value     string    `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}

type Handler struct {
	Tracer oteltrace.Tracer
}

func (h *Handler) ListSamples(res http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	_, span := h.Tracer.Start(ctx, "ListSamples")
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

	sampleJson, err := json.Marshal(samples)
	if err != nil {
		http.Error(res, fmt.Sprintf("Marshalling Error: %v", err), http.StatusInternalServerError)
		span.RecordError(err)
		span.SetStatus(codes.Error, "json.Marhsal(samples)")
		return
	}

	span.SetAttributes(
		attribute.Int("query.result.length", len(samples)),
		attribute.Int("query.result.byte_length", len(sampleJson)),
	)

	res.Header().Set("Content-Type", "application/json")
	res.Write(sampleJson)
}
