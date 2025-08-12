package sample

import (
	"context"
	"net/http"

	"github.com/robert-w/go-server/internal/db/services"
	v1 "github.com/robert-w/go-server/internal/routes/v1"
	"go.opentelemetry.io/otel/codes"
	oteltrace "go.opentelemetry.io/otel/trace"
)

type sampleService interface {
	ListAllSamples(ctx context.Context) (*[]services.Sample, error)
	CreateSamples(ctx context.Context) (*[]services.Sample, error)
	GetSampleById(ctx context.Context) (*services.Sample, error)
	UpdateSampleById(ctx context.Context) (*services.Sample, error)
	DeleteSampleById(ctx context.Context) (*services.Sample, error)
}

type handler struct {
	tracer oteltrace.Tracer
	service sampleService
}

func (h *handler) listSamples(res http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	_, span := h.tracer.Start(ctx, "ListSamples")
	defer span.End()

	samples, serviceErr := h.service.ListAllSamples(ctx)
	if serviceErr != nil {
		span.RecordError(serviceErr)
		span.SetStatus(codes.Error, serviceErr.Error())
	}

	response, err := v1.PrepareResponse(samples, serviceErr)
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
	ctx := req.Context()
	_, span := h.tracer.Start(ctx, "CreateSamples")
	defer span.End()

	samples, serviceErr := h.service.CreateSamples(ctx)
	if serviceErr != nil {
		span.RecordError(serviceErr)
		span.SetStatus(codes.Error, serviceErr.Error())
	}

	response, err := v1.PrepareResponse(samples, serviceErr)
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
	ctx := req.Context()
	_, span := h.tracer.Start(ctx, "ReadSample")
	defer span.End()

	sample, serviceErr := h.service.GetSampleById(ctx)
	if serviceErr != nil {
		span.RecordError(serviceErr)
		span.SetStatus(codes.Error, serviceErr.Error())
	}

	response, err := v1.PrepareResponse(sample, serviceErr)
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
	ctx := req.Context()
	_, span := h.tracer.Start(ctx, "UpdateSample")
	defer span.End()

	sample, serviceErr := h.service.UpdateSampleById(ctx)
	if serviceErr != nil {
		span.RecordError(serviceErr)
		span.SetStatus(codes.Error, serviceErr.Error())
	}

	response, err := v1.PrepareResponse(sample, serviceErr)
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
	ctx := req.Context()
	_, span := h.tracer.Start(ctx, "DeleteSample")
	defer span.End()

	output, serviceErr := h.service.DeleteSampleById(ctx)
	if serviceErr != nil {
		span.RecordError(serviceErr)
		span.SetStatus(codes.Error, serviceErr.Error())
	}

	response, err := v1.PrepareResponse(output, serviceErr)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		http.Error(res, "Error marshalling JSON", http.StatusInternalServerError)
		return
	}

	span.SetStatus(codes.Ok, "Ok")
	res.Header().Set("Content-Type", "application/json")
	res.Write(response)

	span.SetStatus(codes.Ok, "Ok")
	res.WriteHeader(http.StatusNoContent)
}
