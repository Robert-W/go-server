package system

import (
	"net/http/httptest"
	"testing"
)

func TestHealthcheck(t *testing.T) {
	// Mock the req and res
	req := httptest.NewRequest("GET", "http://0.0.0.0:3000/system/health", nil)
	res := httptest.NewRecorder()

	// Invoke the handler with the req and res
	Healthcheck(res, req)

	if res.Code != 200 {
		t.Errorf("Healthcheck must return a 200")
	}
}
