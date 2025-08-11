package sample

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	oteltrace "go.opentelemetry.io/otel/trace"
)

type Handler struct {
	Tracer oteltrace.Tracer
}

func (h *Handler) ListSamples(res http.ResponseWriter, req *http.Request) {
	_, span := h.Tracer.Start(req.Context(), "ListSamples")
	defer span.End()

	samples, err := listSamplesQuery()
	if err != nil {
		http.Error(res, fmt.Sprintf("Error listing all samples: %v", err), http.StatusInternalServerError)
		span.RecordError(err)
		span.SetStatus(codes.Error, "ListSamplesQuery")
		return
	}

	sampleJson, err := json.Marshal(samples)
	if err != nil {
		http.Error(res, fmt.Sprintf("Marshalling Error: %v", err), http.StatusInternalServerError)
		span.RecordError(err)
		span.SetStatus(codes.Error, "json.Marhsal(samples)")
		return
	}

	span.SetAttributes(
		attribute.Int("query.result.length", len(*samples)),
		attribute.Int("query.result.byte_length", len(sampleJson)),
	)

	res.Header().Set("Content-Type", "application/json")
	res.Write(sampleJson)
}

func (h *Handler) CreateSamples(res http.ResponseWriter, req *http.Request) {
	_, span := h.Tracer.Start(req.Context(), "CreateSamples")
	defer span.End()

	samples, err := createSamplesQuery()
	if err != nil {
		http.Error(res, fmt.Sprintf("Error creating samples: %v", err), http.StatusInternalServerError)
		span.RecordError(err)
		span.SetStatus(codes.Error, "createSamplesQuery")
		return
	}

	sampleJson, err := json.Marshal(samples)
	if err != nil {
		http.Error(res, fmt.Sprintf("Marshalling Error: %v", err), http.StatusInternalServerError)
		span.RecordError(err)
		span.SetStatus(codes.Error, "json.Marhsal(samples)")
		return
	}

	span.SetAttributes(
		attribute.Int("query.result.byte_length", len(sampleJson)),
	)

	res.Header().Set("Content-Type", "application/json")
	res.Write(sampleJson)
}

func (h *Handler) ReadSample(res http.ResponseWriter, req *http.Request) {
	_, span := h.Tracer.Start(req.Context(), "ReadSample")
	defer span.End()

	sample, err := readSampleQuery()
	if err != nil {
		http.Error(res, fmt.Sprintf("Error reading sample: %v", err), http.StatusInternalServerError)
		span.RecordError(err)
		span.SetStatus(codes.Error, "readSampleQuery")
		return
	}

	sampleJson, err := json.Marshal(sample)
	if err != nil {
		http.Error(res, fmt.Sprintf("Marshalling Error: %v", err), http.StatusInternalServerError)
		span.RecordError(err)
		span.SetStatus(codes.Error, "json.Marhsal(samples)")
		return
	}

	span.SetAttributes(
		attribute.Int("query.result.byte_length", len(sampleJson)),
	)

	res.Header().Set("Content-Type", "application/json")
	res.Write(sampleJson)
}

func (h *Handler) UpdateSample(res http.ResponseWriter, req *http.Request) {
	_, span := h.Tracer.Start(req.Context(), "UpdateSample")
	defer span.End()

	sample, err := updateSampleQuery()
	if err != nil {
		http.Error(res, fmt.Sprintf("Error updating sample: %v", err), http.StatusInternalServerError)
		span.RecordError(err)
		span.SetStatus(codes.Error, "updateSamplesQuery")
		return
	}

	sampleJson, err := json.Marshal(sample)
	if err != nil {
		http.Error(res, fmt.Sprintf("Marshalling Error: %v", err), http.StatusInternalServerError)
		span.RecordError(err)
		span.SetStatus(codes.Error, "json.Marhsal(sample)")
		return
	}

	span.SetAttributes(
		attribute.Int("query.result.byte_length", len(sampleJson)),
	)

	res.Header().Set("Content-Type", "application/json")
	res.Write(sampleJson)
}

func (h *Handler) DeleteSample(res http.ResponseWriter, req *http.Request) {
	_, span := h.Tracer.Start(req.Context(), "DeleteSample")
	defer span.End()

	vars := mux.Vars(req)
	err := deleteSampleQuery()
	if err != nil {
		http.Error(res, fmt.Sprintf("Error deleting sample: %v", err), http.StatusInternalServerError)
		span.RecordError(err)
		span.SetStatus(codes.Error, "deleteSamplesQuery")
		return
	}

	span.SetAttributes(
		attribute.String("query.deleted.id", vars["id"]),
	)

	res.WriteHeader(http.StatusNoContent)
}
