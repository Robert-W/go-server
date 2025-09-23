package user

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/robert-w/go-server/internal/monitoring"
	"github.com/robert-w/go-server/internal/response"
	"go.opentelemetry.io/otel/codes"
)

type serviceInterface interface {
	list(ctx context.Context) (*[]User, *response.ErrorJsonV1)
	create(ctx context.Context, input *UserPost) (*[]User, *response.ErrorJsonV1)
	get(ctx context.Context, id string) (*User, *response.ErrorJsonV1)
	update(ctx context.Context, id string, input *UserPut) (*User, *response.ErrorJsonV1)
	delete(ctx context.Context, id string) (*User, *response.ErrorJsonV1)
}

type handler struct {
	service serviceInterface
}

func (h *handler) list(res http.ResponseWriter, req *http.Request) {
	ctx, span := monitoring.CreateSpan(req.Context(), "list")
	defer span.End()

	// PrepareResponse won't error as it's just returning the result of
	// json.Marshal on structures we control and are all safe
	users, serviceErr := h.service.list(ctx)
	response, _ := response.NewV1(ctx, users, serviceErr)

	res.Header().Set("Content-Type", "application/json")

	// Set attributes and headers correctly based on what we have in serviceErr
	if serviceErr != nil && serviceErr.StatusCode != 0 {
		res.WriteHeader(serviceErr.StatusCode)
	}

	if serviceErr != nil && serviceErr.Original != nil {
		span.RecordError(serviceErr.Original)
		span.SetStatus(codes.Error, serviceErr.Original.Error())
	}

	res.Write(response)
}

func (h *handler) create(res http.ResponseWriter, req *http.Request) {
	ctx, span := monitoring.CreateSpan(req.Context(), "create")
	defer span.End()

	// PrepareResponse won't error as it's just returning the result of
	// json.Marshal on structures we control and are all safe
	input := req.Context().Value("Input").(*UserPost)
	users, serviceErr := h.service.create(ctx, input)
	response, _ := response.NewV1(ctx, users, serviceErr)

	res.Header().Set("Content-Type", "application/json")

	// Set attributes and headers correctly based on what we have in serviceErr
	if serviceErr != nil && serviceErr.StatusCode != 0 {
		res.WriteHeader(serviceErr.StatusCode)
	}

	if serviceErr != nil && serviceErr.Original != nil {
		span.RecordError(serviceErr.Original)
		span.SetStatus(codes.Error, serviceErr.Original.Error())
	}

	res.Write(response)
}

func (h *handler) get(res http.ResponseWriter, req *http.Request) {
	ctx, span := monitoring.CreateSpan(req.Context(), "get")
	defer span.End()

	// PrepareResponse won't error as it's just returning the result of
	// json.Marshal on structures we control and are all safe
	vars := mux.Vars(req)
	user, serviceErr := h.service.get(ctx, vars["id"])
	response, _ := response.NewV1(ctx, user, serviceErr)

	res.Header().Set("Content-Type", "application/json")

	// Set attributes and headers correctly based on what we have in serviceErr
	if serviceErr != nil && serviceErr.StatusCode != 0 {
		res.WriteHeader(serviceErr.StatusCode)
	}

	if serviceErr != nil && serviceErr.Original != nil {
		span.RecordError(serviceErr.Original)
		span.SetStatus(codes.Error, serviceErr.Original.Error())
	}

	res.Write(response)
}

func (h *handler) update(res http.ResponseWriter, req *http.Request) {
	ctx, span := monitoring.CreateSpan(req.Context(), "update")
	defer span.End()

	// PrepareResponse won't error as it's just returning the result of
	// json.Marshal on structures we control and are all safe
	vars := mux.Vars(req)
	input := req.Context().Value("Input").(*UserPut)
	user, serviceErr := h.service.update(ctx, vars["id"], input)
	response, _ := response.NewV1(ctx, user, serviceErr)

	res.Header().Set("Content-Type", "application/json")

	// Set attributes and headers correctly based on what we have in serviceErr
	if serviceErr != nil && serviceErr.StatusCode != 0 {
		res.WriteHeader(serviceErr.StatusCode)
	}

	if serviceErr != nil && serviceErr.Original != nil {
		span.RecordError(serviceErr.Original)
		span.SetStatus(codes.Error, serviceErr.Original.Error())
	}

	res.Write(response)
}

func (h *handler) delete(res http.ResponseWriter, req *http.Request) {
	ctx, span := monitoring.CreateSpan(req.Context(), "delete")
	defer span.End()

	// PrepareResponse won't error as it's just returning the result of
	// json.Marshal on structures we control and are all safe
	vars := mux.Vars(req)
	output, serviceErr := h.service.delete(ctx, vars["id"])
	response, _ := response.NewV1(ctx, output, serviceErr)

	res.Header().Set("Content-Type", "application/json")

	// Set attributes and headers correctly based on what we have in serviceErr
	if serviceErr != nil && serviceErr.StatusCode != 0 {
		res.WriteHeader(serviceErr.StatusCode)
	}

	if serviceErr != nil && serviceErr.Original != nil {
		span.RecordError(serviceErr.Original)
		span.SetStatus(codes.Error, serviceErr.Original.Error())
	}

	res.Write(response)
}
