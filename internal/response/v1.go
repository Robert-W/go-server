package response

import (
	"context"
	"encoding/json"

	"github.com/noir-cats/go-sample-server/internal/monitoring"
)

type ErrorJsonV1 struct {
	Original   error  `json:"-"`
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
}

type ResponseJsonV1 struct {
	Status string       `json:"status"`
	Result any          `json:"result,omitempty"`
	Error  *ErrorJsonV1 `json:"error,omitempty"`
}

// Wrapper function to take the response from a service, which is either a
// response or nil, or an error or nil. If we have an error, we want to generate
// an error output, otherwise, attempt to generate a success output even if
// response here is null. If this returns an error, its a json.Marshal error
//
// result is any marshallable struct, err is an error interface, errorType is
// referring to error constants from internal/constants/errors.go
func NewV1(ctx context.Context, result any, err *ErrorJsonV1) ([]byte, error) {
	_, span := monitoring.CreateSpan(ctx, "PrepareResponse")
	defer span.End()
	// handle the error scenario first
	if err != nil {
		return json.Marshal(&ResponseJsonV1{Status: "error", Error: err})
	}

	// we have a response, attempt to prepare our output
	return json.Marshal(&ResponseJsonV1{Status: "ok", Result: result})
}
