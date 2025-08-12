package v1

import (
	"encoding/json"
)

type Error struct {
	Original   error  `json:"-"`
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
}

type v1Response struct {
	Status string `json:"status"`
	Result any    `json:"result"`
	Error  *Error `json:"error"`
}

// Wrapper function to take the response from a service, which is either a
// response or nil, or an error or nil. If we have an error, we want to generate
// an error output, otherwise, attempt to generate a success output even if
// response here is null. If this returns an error, its a json.Marshal error
//
// result is any marshallable struct, err is an error interface, errorType is
// referring to error constants from internal/constants/errors.go
func PrepareResponse(result any, err *Error) ([]byte, error) {
	// handle the error scenario first
	if err != nil {
		return json.Marshal(&v1Response{Status: "error", Error: err})
	}

	// we have a response, attempt to prepare our output
	return json.Marshal(&v1Response{Status: "ok", Result: result})
}
