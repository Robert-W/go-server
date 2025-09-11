package middleware

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/robert-w/go-server/internal/monitoring"
	"github.com/robert-w/go-server/internal/requestutil"
	"go.opentelemetry.io/otel/codes"
)

// This middleware should parse the request into a struct and perform validation
// on that struct. If the binding/validation is successful, attach the bound
// struct to the request context so it can be retrieved later. In failure
// scenarios, it should do the following:
// - Unable to Unmarshal Json into the given struct - 400
// - Failure to validate the struct                 - 422

// Usage:
// subrouter.Handle("/users", ValidateMiddleware(validator, &User{}, userHandler.list)).Methods("GET")
func ValidateMiddleware(
	utils *requestutil.RequestUtils,
	structPtr any,
	next http.HandlerFunc) func(w http.ResponseWriter, r *http.Request) {

	decoder := utils.SchemaDecoder
	validate := utils.Validate

	return func(w http.ResponseWriter, req *http.Request) {
		ctx, span := monitoring.CreateSpan(req.Context(), "ValidateMiddleware")
		defer span.End()

		// Return from any error so we don't invoke the next handler

		// ParseForm will attempt to get form data and/or query parameters, can add
		// more type later
		if req.URL.RawQuery != "" || req.Header.Get("Content-Type") == "application/x-www-form-urlencoded" {
			if err := req.ParseForm(); err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())

				slog.Error("Error from parseForm", "error", err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			if err := decoder.Decode(structPtr, req.Form); err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())

				slog.Error("Error decoding req.Form", "error", err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}
		}

		// Make an attempt to bind the req.Body to our provided structPtr
		if req.ContentLength > 0 {
			jsondecoder := json.NewDecoder(req.Body)
			if err := jsondecoder.Decode(structPtr); err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())

				slog.Error("Error decoding req.Body", "error", err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}
		}

		err := validate.Struct(structPtr)
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())

			var validateErrs validator.ValidationErrors
			var details []string
			if errors.As(err, &validateErrs) {
				for _, e := range validateErrs {
					details = append(
						details,
						fmt.Sprintf("Field: %s, Value: %s", e.StructField(), e.Value()),
					)
				}
			}

			slog.Error("Error validating the struct", "error(s)", err, "details", details)
			w.WriteHeader(http.StatusUnprocessableEntity)
			return
		}

		// Store the bound struct in the request context
		ctx = context.WithValue(ctx, "Input", structPtr)
		next(w, req.WithContext(ctx))
	}
}
