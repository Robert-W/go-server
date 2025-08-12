package sample

import (
	"net/http"

	"github.com/robert-w/go-server/internal/monitoring"
	v1 "github.com/robert-w/go-server/internal/routes/v1"
	"go.opentelemetry.io/otel/codes"
)

type handler struct {
	service *sampleService
}

func (h *handler) listSamples(res http.ResponseWriter, req *http.Request) {
	ctx, span := monitoring.CreateSpan(req.Context(), "listSamples")
	defer span.End()

	samples, serviceErr := h.service.listAllSamples(ctx)
	if serviceErr != nil {
		span.RecordError(serviceErr.Original)
		span.SetStatus(codes.Error, serviceErr.Original.Error())
	}

	response, err := v1.PrepareResponse(ctx, samples, serviceErr)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		http.Error(res, "Error marshalling JSON", http.StatusInternalServerError)
		return
	}

	span.SetStatus(codes.Ok, "Ok")
	res.Header().Set("Content-Type", "application/json")
	res.Write(response)
}

func (h *handler) createSamples(res http.ResponseWriter, req *http.Request) {
	ctx, span := monitoring.CreateSpan(req.Context(), "createSamples")
	defer span.End()

	samples, serviceErr := h.service.createSamples(ctx)
	if serviceErr != nil {
		span.RecordError(serviceErr.Original)
		span.SetStatus(codes.Error, serviceErr.Original.Error())
	}

	response, err := v1.PrepareResponse(ctx, samples, serviceErr)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		http.Error(res, "Error marshalling JSON", http.StatusInternalServerError)
		return
	}

	span.SetStatus(codes.Ok, "Ok")
	res.Header().Set("Content-Type", "application/json")
	res.Write(response)
}

func (h *handler) readSample(res http.ResponseWriter, req *http.Request) {
	ctx, span := monitoring.CreateSpan(req.Context(), "readSample")
	defer span.End()

	sample, serviceErr := h.service.getSampleById(ctx)
	if serviceErr != nil {
		span.RecordError(serviceErr.Original)
		span.SetStatus(codes.Error, serviceErr.Original.Error())
	}

	response, err := v1.PrepareResponse(ctx, sample, serviceErr)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		http.Error(res, "Error marshalling JSON", http.StatusInternalServerError)
		return
	}

	span.SetStatus(codes.Ok, "Ok")
	res.Header().Set("Content-Type", "application/json")
	res.Write(response)
}

func (h *handler) updateSample(res http.ResponseWriter, req *http.Request) {
	ctx, span := monitoring.CreateSpan(req.Context(), "updateSample")
	defer span.End()

	sample, serviceErr := h.service.updateSampleById(ctx)
	if serviceErr != nil {
		span.RecordError(serviceErr.Original)
		span.SetStatus(codes.Error, serviceErr.Original.Error())
	}

	response, err := v1.PrepareResponse(ctx, sample, serviceErr)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		http.Error(res, "Error marshalling JSON", http.StatusInternalServerError)
		return
	}

	span.SetStatus(codes.Ok, "Ok")
	res.Header().Set("Content-Type", "application/json")
	res.Write(response)
}

func (h *handler) deleteSample(res http.ResponseWriter, req *http.Request) {
	ctx, span := monitoring.CreateSpan(req.Context(), "deleteSample")
	defer span.End()

	output, serviceErr := h.service.deleteSampleById(ctx)
	if serviceErr != nil {
		span.RecordError(serviceErr.Original)
		span.SetStatus(codes.Error, serviceErr.Original.Error())
	}

	response, err := v1.PrepareResponse(ctx, output, serviceErr)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		http.Error(res, "Error marshalling JSON", http.StatusInternalServerError)
		return
	}

	span.SetStatus(codes.Ok, "Ok")
	res.Header().Set("Content-Type", "application/json")
	res.Write(response)
}
