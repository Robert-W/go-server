package middleware

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/noir-cats/go-sample-server/internal/monitoring"
	"github.com/noir-cats/go-sample-server/internal/requestutil"
	"github.com/noir-cats/go-sample-server/internal/response"
	"go.opentelemetry.io/otel/codes"
)

// This middleware should parse the request into a struct and perform validation
// on that struct. If the binding/validation is successful, attach the bound
// struct to the request context so it can be retrieved later. In failure
// scenarios, it should do the following:
// - Unable to Unmarshal Json into the given struct - 400
// - Failure to validate the struct                 - 422
//
// This function currently returns a NewV1 response type since this server only
// has one version. Future versions can abstract that out when needed.
//
// Usage:
// subrouter.Handle("/users", ValidateMiddleware(validator, &User{}, userHandler.list)).Methods("GET")
func ValidateMiddleware(
	utils *requestutil.RequestUtils,
	structPtr any,
	next http.HandlerFunc) func(http.ResponseWriter, *http.Request) {

	decoder := utils.SchemaDecoder
	validate := utils.Validate

	return func(w http.ResponseWriter, r *http.Request) {
		ctx, span := monitoring.CreateSpan(r.Context(), "ValidateMiddleware")
		defer span.End()

		// Return from any error so we don't invoke the next handler

		// ParseForm will attempt to get form data and/or query parameters, can add
		// more type later
		if r.URL.RawQuery != "" || r.Header.Get("Content-Type") == "application/x-www-form-urlencoded" {
			if err := r.ParseForm(); err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())

				errjson := &response.ErrorJsonV1{
					Message:    "Error parsing form and/or query parameters",
					Original:   err,
					StatusCode: http.StatusBadRequest,
				}

				// Ignore the error here, this won't cause a marshalling issue
				res, _ := response.NewV1(ctx, nil, errjson)
				w.WriteHeader(errjson.StatusCode)
				w.Write(res)
				return
			}

			if err := decoder.Decode(structPtr, r.Form); err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())

				errjson := &response.ErrorJsonV1{
					Message:    "Error decoding form into desired type",
					Original:   err,
					StatusCode: http.StatusBadRequest,
				}

				// Ignore the error here, this won't cause a marshalling issue
				res, _ := response.NewV1(ctx, nil, errjson)
				w.WriteHeader(errjson.StatusCode)
				w.Write(res)
				return
			}
		}

		// Make an attempt to bind the req.Body to our provided structPtr
		if r.ContentLength > 0 {
			jsondecoder := json.NewDecoder(r.Body)
			if err := jsondecoder.Decode(structPtr); err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())

				errjson := &response.ErrorJsonV1{
					Message:    "Error decoding request body into desired type",
					Original:   err,
					StatusCode: http.StatusBadRequest,
				}

				// Ignore the error here, this won't cause a marshalling issue
				res, _ := response.NewV1(ctx, nil, errjson)
				w.WriteHeader(errjson.StatusCode)
				w.Write(res)
				return
			}
		}

		err := validate.Struct(structPtr)
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())

			// If we want more custom error messaging, the validator library does have
			// some options for getting more details out, but this is fine for now
			errjson := &response.ErrorJsonV1{
				Message:    err.Error(),
				Original:   err,
				StatusCode: http.StatusUnprocessableEntity,
			}

			// Ignore the error here, this won't cause a marshalling issue
			res, _ := response.NewV1(ctx, nil, errjson)
			w.WriteHeader(errjson.StatusCode)
			w.Write(res)
			return
		}

		// Store the bound struct in the request context
		ctx = context.WithValue(ctx, "Input", structPtr)
		next(w, r.WithContext(ctx))
	}
}
