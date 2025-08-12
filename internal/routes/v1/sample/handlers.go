package sample

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/robert-w/go-server/internal/db/services"
	"go.opentelemetry.io/otel/codes"
	oteltrace "go.opentelemetry.io/otel/trace"
)

type handler struct {
	tracer oteltrace.Tracer
	service *services.SampleService
}

func (h *handler) listSamples(res http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	_, span := h.tracer.Start(ctx, "ListSamples")
	defer span.End()

	samples, err := h.service.ListAllSamples(ctx)
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
	res.Header().Set("Content-Type", "application/json")
	res.Write(sampleJson)
}

func (h *handler) createSamples(res http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	_, span := h.tracer.Start(ctx, "CreateSamples")
	defer span.End()

	samples, err := h.service.CreateSamples(ctx)
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
	res.Header().Set("Content-Type", "application/json")
	res.Write(sampleJson)
}

func (h *handler) readSample(res http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	_, span := h.tracer.Start(ctx, "ReadSample")
	defer span.End()

	sample, err := h.service.GetSampleById(ctx)
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
	res.Header().Set("Content-Type", "application/json")
	res.Write(sampleJson)
}

func (h *handler) updateSample(res http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	_, span := h.tracer.Start(ctx, "UpdateSample")
	defer span.End()

	sample, err := h.service.UpdateSampleById(ctx)
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
	res.Header().Set("Content-Type", "application/json")
	res.Write(sampleJson)
}

func (h *handler) deleteSample(res http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	_, span := h.tracer.Start(ctx, "DeleteSample")
	defer span.End()

	err := h.service.DeleteSampleById(ctx)
	if err != nil {
		http.Error(res, fmt.Sprintf("Error deleting sample: %v", err), http.StatusInternalServerError)
		span.RecordError(err)
		span.SetStatus(codes.Error, "deleteSamplesQuery")
		return
	}

	span.SetStatus(codes.Ok, "Ok")
	res.WriteHeader(http.StatusNoContent)
}
