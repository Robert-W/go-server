package v1

import (
	"context"
	"errors"
	"testing"
)

type response struct {
	Foo string `json:"foo"`
}

type badResponse struct {
	Foo chan int
}

func TestPrepareResponse(t *testing.T) {
	ctx := context.Background()

	t.Run("should return a valid success response when no error given", func(t *testing.T) {
		goodResponse := &response{Foo: "Scooby"}
		output, err := PrepareResponse(ctx, goodResponse, nil)

		if err != nil {
			t.Errorf("Received an marshalling error when we should have a v1Response: %v", err)
		}

		expected := `{"status":"ok","result":{"foo":"Scooby"}}`
		actual := string(output)
		if string(actual) != expected {
			t.Errorf("Expected: %v, Actual: %v", expected, actual)
		}
	})

	t.Run("should return a valid error response when given an error", func(t *testing.T) {
		goodErr := &Error{
			Original:   errors.New("OG"),
			Message:    "Scooby",
			StatusCode: 500,
		}
		output, err := PrepareResponse(ctx, nil, goodErr)

		if err != nil {
			t.Errorf("Received an marshalling error when we should have a v1Response: %v", err)
		}

		expected := `{"status":"error","error":{"message":"Scooby","statusCode":500}}`
		actual := string(output)
		if string(actual) != expected {
			t.Errorf("Expected: %v, Actual: %v", expected, actual)
		}
	})

	// If someone passes in a struct that contains fields that cannot be
	// marshalled, this should pass the marshalling error back, don't test on the
	// message here in case it changes, just that we get back what we expect
	t.Run("should return a marshalling error when the provided response is not marshallable", func(t *testing.T) {
		badResponse := &badResponse{Foo: make(chan int)}
		output, err := PrepareResponse(ctx, badResponse, nil)

		if output != nil {
			t.Errorf("Bad Response should not be successfully serialized")
		}

		if err == nil {
			t.Error("We should get a marshalling error")
		}
	})

}
