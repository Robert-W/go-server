package v1

import (
	"encoding/json"
)

type Error struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

type v1Response struct {
	Status string `json:"status"`
	Result any    `json:"result"`
	Error  string `json:"error"`
}

// Wrapper function to take the response from a service, which is either a
// response or nil, or an error or nil. If we have an error, we want to generate
// an error output, otherwise, attempt to generate a success output even if
// response here is null. If this returns an error, its a json.Marshal error
//
// result is any marshallable struct, err is an error interface, errorType is
// referring to error constants from internal/constants/errors.go
func PrepareResponse(result any, err error) ([]byte, error) {
	// handle the error scenario first
	if err != nil {
		return json.Marshal(&v1Response{
			Status: "error",
			Error:  err.Error(),
		})
	}

	// we have a response, attempt to prepare our output
	return json.Marshal(&v1Response{Status: "ok", Result: result})
}
