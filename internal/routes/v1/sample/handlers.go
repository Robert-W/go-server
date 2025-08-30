package sample

import (
	"context"
	"net/http"

	"github.com/robert-w/go-server/internal/monitoring"
	v1 "github.com/robert-w/go-server/internal/routes/v1"
	"go.opentelemetry.io/otel/codes"
)

type serviceInterface interface {
	listAllSamples(ctx context.Context) (*[]sample, *v1.Error)
	createSamples(ctx context.Context) (*[]sample, *v1.Error)
	getSampleById(ctx context.Context) (*sample, *v1.Error)
	updateSampleById(ctx context.Context) (*sample, *v1.Error)
	deleteSampleById(ctx context.Context) (*sample, *v1.Error)
}

type handler struct {
	service serviceInterface
}

func (h *handler) list(res http.ResponseWriter, req *http.Request) {
	ctx, span := monitoring.CreateSpan(req.Context(), "list")
	defer span.End()

	// PrepareResponse won't error as it's just returning the result of
	// json.Marshal on structures we control and are all safe
	samples, serviceErr := h.service.listAllSamples(ctx)
	response, _ := v1.PrepareResponse(ctx, samples, serviceErr)

	// Set attributes and headers correctly based on what we have in serviceErr
	if serviceErr != nil && serviceErr.StatusCode != 0 {
		res.WriteHeader(serviceErr.StatusCode)
	}

	if serviceErr != nil && serviceErr.Original != nil {
		span.RecordError(serviceErr.Original)
		span.SetStatus(codes.Error, serviceErr.Original.Error())
	} else {
		span.SetStatus(codes.Ok, "Ok")
	}

	res.Header().Set("Content-Type", "application/json")
	res.Write(response)
}

func (h *handler) create(res http.ResponseWriter, req *http.Request) {
	ctx, span := monitoring.CreateSpan(req.Context(), "create")
	defer span.End()

	// PrepareResponse won't error as it's just returning the result of
	// json.Marshal on structures we control and are all safe
	samples, serviceErr := h.service.createSamples(ctx)
	response, _ := v1.PrepareResponse(ctx, samples, serviceErr)

	// Set attributes and headers correctly based on what we have in serviceErr
	if serviceErr != nil && serviceErr.StatusCode != 0 {
		res.WriteHeader(serviceErr.StatusCode)
	}

	if serviceErr != nil && serviceErr.Original != nil {
		span.RecordError(serviceErr.Original)
		span.SetStatus(codes.Error, serviceErr.Original.Error())
	} else {
		span.SetStatus(codes.Ok, "Ok")
	}

	res.Header().Set("Content-Type", "application/json")
	res.Write(response)
}

func (h *handler) get(res http.ResponseWriter, req *http.Request) {
	ctx, span := monitoring.CreateSpan(req.Context(), "get")
	defer span.End()

	// PrepareResponse won't error as it's just returning the result of
	// json.Marshal on structures we control and are all safe
	sample, serviceErr := h.service.getSampleById(ctx)
	response, _ := v1.PrepareResponse(ctx, sample, serviceErr)

	// Set attributes and headers correctly based on what we have in serviceErr
	if serviceErr != nil && serviceErr.StatusCode != 0 {
		res.WriteHeader(serviceErr.StatusCode)
	}

	if serviceErr != nil && serviceErr.Original != nil {
		span.RecordError(serviceErr.Original)
		span.SetStatus(codes.Error, serviceErr.Original.Error())
	} else {
		span.SetStatus(codes.Ok, "Ok")
	}

	res.Header().Set("Content-Type", "application/json")
	res.Write(response)
}

func (h *handler) update(res http.ResponseWriter, req *http.Request) {
	ctx, span := monitoring.CreateSpan(req.Context(), "update")
	defer span.End()

	// PrepareResponse won't error as it's just returning the result of
	// json.Marshal on structures we control and are all safe
	sample, serviceErr := h.service.updateSampleById(ctx)
	response, _ := v1.PrepareResponse(ctx, sample, serviceErr)

	// Set attributes and headers correctly based on what we have in serviceErr
	if serviceErr != nil && serviceErr.StatusCode != 0 {
		res.WriteHeader(serviceErr.StatusCode)
	}

	if serviceErr != nil && serviceErr.Original != nil {
		span.RecordError(serviceErr.Original)
		span.SetStatus(codes.Error, serviceErr.Original.Error())
	} else {
		span.SetStatus(codes.Ok, "Ok")
	}

	res.Header().Set("Content-Type", "application/json")
	res.Write(response)
}

func (h *handler) delete(res http.ResponseWriter, req *http.Request) {
	ctx, span := monitoring.CreateSpan(req.Context(), "delete")
	defer span.End()

	// PrepareResponse won't error as it's just returning the result of
	// json.Marshal on structures we control and are all safe
	output, serviceErr := h.service.deleteSampleById(ctx)
	response, _ := v1.PrepareResponse(ctx, output, serviceErr)

	// Set attributes and headers correctly based on what we have in serviceErr
	if serviceErr != nil && serviceErr.StatusCode != 0 {
		res.WriteHeader(serviceErr.StatusCode)
	}

	if serviceErr != nil && serviceErr.Original != nil {
		span.RecordError(serviceErr.Original)
		span.SetStatus(codes.Error, serviceErr.Original.Error())
	} else {
		span.SetStatus(codes.Ok, "Ok")
	}

	res.Header().Set("Content-Type", "application/json")
	res.Write(response)
}
