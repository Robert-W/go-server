package sample

import (
	"context"
	"encoding/json"
	"net/http/httptest"
	"testing"
	"time"

	v1 "github.com/robert-w/go-server/internal/routes/v1"
)

// Create a mock that returns a successful response
type mockSampleService struct{}

func (m *mockSampleService) listAllSamples(ctx context.Context) (*[]sample, *v1.Error) {
	samples := []sample{{Id: "111", Value: "First Sample", Timestamp: time.Now()}}
	return &samples, nil
}

func (m *mockSampleService) createSamples(ctx context.Context) (*[]sample, *v1.Error) {
	samples := []sample{{Id: "111", Value: "First Sample", Timestamp: time.Now()}}
	return &samples, nil
}

func (m *mockSampleService) getSampleById(ctx context.Context) (*sample, *v1.Error) {
	return &sample{Id: "111", Value: "Sample"}, nil
}

func (m *mockSampleService) updateSampleById(ctx context.Context) (*sample, *v1.Error) {
	return &sample{Id: "111", Value: "Sample"}, nil
}

func (m *mockSampleService) deleteSampleById(ctx context.Context) (*sample, *v1.Error) {
	return &sample{Id: "111", Value: "Sample"}, nil
}

// Create a mock that returns a versioned error
type mockSampleServiceErr struct{}

func (m *mockSampleServiceErr) listAllSamples(ctx context.Context) (*[]sample, *v1.Error) {
	return nil, &v1.Error{Message: "Scooby Dooby Doo", StatusCode: 500}
}

func (m *mockSampleServiceErr) createSamples(ctx context.Context) (*[]sample, *v1.Error) {
	return nil, &v1.Error{Message: "Scooby Dooby Doo", StatusCode: 500}
}

func (m *mockSampleServiceErr) getSampleById(ctx context.Context) (*sample, *v1.Error) {
	return nil, &v1.Error{Message: "Scooby Dooby Doo", StatusCode: 404}
}

func (m *mockSampleServiceErr) updateSampleById(ctx context.Context) (*sample, *v1.Error) {
	return nil, &v1.Error{Message: "Scooby Dooby Doo", StatusCode: 404}
}

func (m *mockSampleServiceErr) deleteSampleById(ctx context.Context) (*sample, *v1.Error) {
	return nil, &v1.Error{Message: "Scooby Dooby Doo", StatusCode: 404}
}

// Types for parsing responses
type mockResultSampleList struct {
	Result []sample `json:"result"`
}

type mockResultSample struct {
	Result sample `json:"result"`
}

type mockResultV1Error struct {
	Error v1.Error `json:"error"`
}

func TestListSamples(t *testing.T) {
	testHandler := handler{service: &mockSampleService{}}
	testHandlerErr := handler{service: &mockSampleServiceErr{}}

	t.Run("should return samples in the format of a v1Response", func(t *testing.T) {
		req := httptest.NewRequest("GET", "http://0.0.0.0:3000/api/v1/samples", nil)
		res := httptest.NewRecorder()

		testHandler.listSamples(res, req)

		if res.Code != 200 {
			t.Error("listSamples should return a 200")
		}

		var result mockResultSampleList
		err := json.Unmarshal(res.Body.Bytes(), &result)

		if err != nil {
			t.Errorf("Unable to decode response: %v", err)
		}

		if len(result.Result) != 1 {
			t.Error("Result is not the correct length")
		}

		if result.Result[0].Id != "111" {
			t.Error("Result does not have the correct ID")
		}
	})

	t.Run("should return a v1Error if the underlying service returns an error", func(t *testing.T) {
		req := httptest.NewRequest("GET", "http://0.0.0.0:3000/api/v1/samples", nil)
		res := httptest.NewRecorder()

		testHandlerErr.listSamples(res, req)

		if res.Code != 500 {
			t.Errorf("listSamples should return a 500, got %d", res.Code)
		}

		var result mockResultV1Error
		err := json.Unmarshal(res.Body.Bytes(), &result)

		if err != nil {
			t.Errorf("Unable to decode response: %v", err)
		}

		if result.Error.Message != "Scooby Dooby Doo" {
			t.Error("Did not receive the expected error message")
		}
	})
}

func TestCreateSamples(t *testing.T) {
	testHandler := handler{service: &mockSampleService{}}
	testHandlerErr := handler{service: &mockSampleServiceErr{}}

	t.Run("should return the created samples in the format of a v1Response", func(t *testing.T) {
		req := httptest.NewRequest("POST", "http://0.0.0.0:3000/api/v1/samples", nil)
		res := httptest.NewRecorder()

		testHandler.createSamples(res, req)

		if res.Code != 200 {
			t.Error("createSamples should return a 200")
		}

		var result mockResultSampleList
		err := json.Unmarshal(res.Body.Bytes(), &result)

		if err != nil {
			t.Errorf("Unable to decode response: %v", err)
		}

		if len(result.Result) != 1 {
			t.Error("Result is not the correct length")
		}

		if result.Result[0].Id != "111" {
			t.Error("Result does not have the correct ID")
		}
	})

	t.Run("should return a v1Error if the underlying service returns an error", func(t *testing.T) {
		req := httptest.NewRequest("POST", "http://0.0.0.0:3000/api/v1/samples", nil)
		res := httptest.NewRecorder()

		testHandlerErr.createSamples(res, req)

		if res.Code != 500 {
			t.Errorf("createSamples should return a 500, got %d", res.Code)
		}

		var result mockResultV1Error
		err := json.Unmarshal(res.Body.Bytes(), &result)

		if err != nil {
			t.Errorf("Unable to decode response: %v", err)
		}

		if result.Error.Message != "Scooby Dooby Doo" {
			t.Error("Did not receive the expected error message")
		}
	})
}

func TestReadSample(t *testing.T) {
	testHandler := handler{service: &mockSampleService{}}
	testHandlerErr := handler{service: &mockSampleServiceErr{}}

	t.Run("should return the sample in the format of a v1Response", func(t *testing.T) {
		req := httptest.NewRequest("GET", "http://0.0.0.0:3000/api/v1/samples/111", nil)
		res := httptest.NewRecorder()

		testHandler.readSample(res, req)

		if res.Code != 200 {
			t.Error("readSample should return a 200")
		}

		var result mockResultSample
		err := json.Unmarshal(res.Body.Bytes(), &result)

		if err != nil {
			t.Errorf("Unable to decode response: %v", err)
		}

		if result.Result.Id != "111" {
			t.Error("Result does not have the correct ID")
		}
	})

	t.Run("should return a v1Error if the underlying service returns an error", func(t *testing.T) {
		req := httptest.NewRequest("GET", "http://0.0.0.0:3000/api/v1/samples/111", nil)
		res := httptest.NewRecorder()

		testHandlerErr.readSample(res, req)

		if res.Code != 404 {
			t.Errorf("readSample should return a 404, got %d", res.Code)
		}

		var result mockResultV1Error
		err := json.Unmarshal(res.Body.Bytes(), &result)

		if err != nil {
			t.Errorf("Unable to decode response: %v", err)
		}

		if result.Error.Message != "Scooby Dooby Doo" {
			t.Error("Did not receive the expected error message")
		}
	})
}

func TestUpdateSample(t *testing.T) {
	testHandler := handler{service: &mockSampleService{}}
	testHandlerErr := handler{service: &mockSampleServiceErr{}}

	t.Run("should return the updated sample in the format of a v1Response", func(t *testing.T) {
		req := httptest.NewRequest("PUT", "http://0.0.0.0:3000/api/v1/samples/111", nil)
		res := httptest.NewRecorder()

		testHandler.updateSample(res, req)

		if res.Code != 200 {
			t.Error("updateSample should return a 200")
		}

		var result mockResultSample
		err := json.Unmarshal(res.Body.Bytes(), &result)

		if err != nil {
			t.Errorf("Unable to decode response: %v", err)
		}

		if result.Result.Id != "111" {
			t.Error("Result does not have the correct ID")
		}
	})

	t.Run("should return a v1Error if the underlying service returns an error", func(t *testing.T) {
		req := httptest.NewRequest("PUT", "http://0.0.0.0:3000/api/v1/samples/111", nil)
		res := httptest.NewRecorder()

		testHandlerErr.updateSample(res, req)

		if res.Code != 404 {
			t.Errorf("updateSample should return a 404, got %d", res.Code)
		}

		var result mockResultV1Error
		err := json.Unmarshal(res.Body.Bytes(), &result)

		if err != nil {
			t.Errorf("Unable to decode response: %v", err)
		}

		if result.Error.Message != "Scooby Dooby Doo" {
			t.Error("Did not receive the expected error message")
		}
	})
}

func TestDeleteSample(t *testing.T) {
	testHandler := handler{service: &mockSampleService{}}
	testHandlerErr := handler{service: &mockSampleServiceErr{}}

	t.Run("should return the id of the deleted sample in the format of a v1Response", func(t *testing.T) {
		req := httptest.NewRequest("DELETE", "http://0.0.0.0:3000/api/v1/samples/111", nil)
		res := httptest.NewRecorder()

		testHandler.deleteSample(res, req)

		if res.Code != 200 {
			t.Error("deleteSample should return a 200")
		}

		var result mockResultSample
		err := json.Unmarshal(res.Body.Bytes(), &result)

		if err != nil {
			t.Errorf("Unable to decode response: %v", err)
		}

		if result.Result.Id != "111" {
			t.Error("Result does not have the correct ID")
		}
	})

	t.Run("should return a v1Error if the underlying service returns an error", func(t *testing.T) {
		req := httptest.NewRequest("DELETE", "http://0.0.0.0:3000/api/v1/samples/111", nil)
		res := httptest.NewRecorder()

		testHandlerErr.deleteSample(res, req)

		if res.Code != 404 {
			t.Errorf("deleteSample should return a 404, got %d", res.Code)
		}

		var result mockResultV1Error
		err := json.Unmarshal(res.Body.Bytes(), &result)

		if err != nil {
			t.Errorf("Unable to decode response: %v", err)
		}

		if result.Error.Message != "Scooby Dooby Doo" {
			t.Error("Did not receive the expected error message")
		}
	})
}
