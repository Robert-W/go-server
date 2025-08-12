package sample

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/robert-w/go-server/internal/db/services"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	oteltrace "go.opentelemetry.io/otel/trace"
)

type Handler struct {
	Tracer oteltrace.Tracer
	Service *services.SampleService
}

func (h *Handler) ListSamples(res http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	_, span := h.Tracer.Start(ctx, "ListSamples")
	defer span.End()

	samples, err := h.Service.ListAllSamples(ctx)
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

	span.SetStatus(codes.Ok, "Ok")
	span.SetAttributes(
		attribute.Int("query.result.length", len(*samples)),
		attribute.Int("query.result.byte_length", len(sampleJson)),
	)

	res.Header().Set("Content-Type", "application/json")
	res.Write(sampleJson)
}

func (h *Handler) CreateSamples(res http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	_, span := h.Tracer.Start(ctx, "CreateSamples")
	defer span.End()

	samples, err := h.Service.CreateSamples(ctx)
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

	span.SetStatus(codes.Ok, "Ok")
	span.SetAttributes(
		attribute.Int("query.result.byte_length", len(sampleJson)),
	)

	res.Header().Set("Content-Type", "application/json")
	res.Write(sampleJson)
}

func (h *Handler) ReadSample(res http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	_, span := h.Tracer.Start(ctx, "ReadSample")
	defer span.End()

	sample, err := h.Service.GetSampleById(ctx)
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

	span.SetStatus(codes.Ok, "Ok")
	span.SetAttributes(
		attribute.Int("query.result.byte_length", len(sampleJson)),
	)

	res.Header().Set("Content-Type", "application/json")
	res.Write(sampleJson)
}

func (h *Handler) UpdateSample(res http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	_, span := h.Tracer.Start(ctx, "UpdateSample")
	defer span.End()

	sample, err := h.Service.UpdateSampleById(ctx)
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

	span.SetStatus(codes.Ok, "Ok")
	span.SetAttributes(
		attribute.Int("query.result.byte_length", len(sampleJson)),
	)

	res.Header().Set("Content-Type", "application/json")
	res.Write(sampleJson)
}

func (h *Handler) DeleteSample(res http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	_, span := h.Tracer.Start(ctx, "DeleteSample")
	defer span.End()

	vars := mux.Vars(req)
	err := h.Service.DeleteSampleById(ctx)
	if err != nil {
		http.Error(res, fmt.Sprintf("Error deleting sample: %v", err), http.StatusInternalServerError)
		span.RecordError(err)
		span.SetStatus(codes.Error, "deleteSamplesQuery")
		return
	}

	span.SetStatus(codes.Ok, "Ok")
	span.SetAttributes(
		attribute.String("query.deleted.id", vars["id"]),
	)

	res.WriteHeader(http.StatusNoContent)
}
