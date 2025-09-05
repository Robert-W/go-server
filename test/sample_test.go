package test

import (
	"fmt"
	"net/http"
	"testing"
)

// This file is for performing integration tests on the v1 sample endpoints
type SampleResponseMany struct {
	Status string   `json:"status"`
	Result []Sample `json:"result"`
}

type SampleResponseSingle struct {
	Status string `json:"status"`
	Result Sample `json:"result"`
}

type Sample struct {
	Id string `json:"id"`
}

func TestSampleEndpoints(t *testing.T) {
	var err error
	var response *http.Response
	var request *http.Request
	var url string

	client := http.Client{}

	// Create a sample
	response, err = http.Post("0.0.0.0:3000/api/v1/samples", "application/json", nil)
	if err != nil {
		t.Errorf("Failed to create samples: %v", err)
	}

	println("Create Response: %v", response)

	// Update that sample
	url = fmt.Sprintf("0.0.0.0:3000/api/v1/samples/%s", "<id>")
	request, err = http.NewRequest(http.MethodPut, url, nil)
	if err != nil {
		t.Errorf("Failed to create put request: %v", err)
	}

	response, err = client.Do(request)
	if err != nil {
		t.Errorf("Failed to update sample: %v", err)
	}

	println("Put Response: %v", response)

	// Get the sample by id
	url = fmt.Sprintf("0.0.0.0:3000/api/v1/samples/%s", "<id>")
	response, err = http.Get(url)
	if err != nil {
		t.Errorf("Failed to get sample: %v", err)
	}

	println("Get Response: %v", response)

	// List all samples
	response, err = http.Get("0.0.0.0:3000/api/v1/samples")
	if err != nil {
		t.Errorf("Failed to list samples: %v", err)
	}

	println("List Response: %v", response)

	// Delete a sample
	url = fmt.Sprintf("0.0.0.0:3000/api/v1/samples/%s", "<id>")
	request, err = http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		t.Errorf("Failed to create delete request: %v", err)
	}

	response, err = client.Do(request)
	if err != nil {
		t.Errorf("Failed to delete sample: %v", err)
	}

	println("Delete Response: %v", response)
}
