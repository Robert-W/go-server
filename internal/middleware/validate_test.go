package middleware

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/noir-cats/go-sample-server/internal/requestutil"
)

type Thing struct {
	Email string `json:"email" validate:"required,email"`
}

type ThingNoValidation struct {
	Email string `json:"email"`
}

func TestValidateMiddleware(t *testing.T) {
	utils := requestutil.New(
		requestutil.WithSchemaDecoder(),
		requestutil.WithValidator(),
	)

	t.Run("should return 400 if unable to unmarhsal the request", func(t *testing.T) {
		next := func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}

		// Make some invalid json
		reader := strings.NewReader(`{"email": nil}`)
		req := httptest.NewRequest("POST", "/", reader)
		req.Header.Add("content-type", "application/json")
		res := httptest.NewRecorder()
		middleware := ValidateMiddleware(utils, &Thing{}, next)

		middleware(res, req)

		if res.Code != 400 {
			t.Errorf("expected 400, got %d", res.Code)
		}
	})

	t.Run("should return 422 if it fails the validation", func(t *testing.T) {
		// This should not be called because the middleware should error out
		next := func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}

		reader := strings.NewReader(`{"email": "scooby"}`)
		req := httptest.NewRequest("POST", "/", reader)
		req.Header.Add("content-type", "application/json")
		res := httptest.NewRecorder()
		middleware := ValidateMiddleware(utils, &Thing{}, next)

		middleware(res, req)

		if res.Code != 422 {
			t.Errorf("expected 422, got %d", res.Code)
		}
	})

	t.Run("should parse the request body and insert the result into the request context", func(t *testing.T) {
		email := "scooby@dooby.doo"
		next := func(w http.ResponseWriter, r *http.Request) {
			input := r.Context().Value("Input").(*Thing)

			if input.Email != email {
				t.Errorf("Thing was not successfully parsed")
			}

			w.WriteHeader(http.StatusOK)
		}

		reader := strings.NewReader(fmt.Sprintf(`{"email": "%s"}`, email))
		req := httptest.NewRequest("POST", "/", reader)
		req.Header.Add("content-type", "application/json")
		res := httptest.NewRecorder()
		middleware := ValidateMiddleware(utils, &Thing{}, next)

		middleware(res, req)

		if res.Code != 200 {
			t.Errorf("expected 200, got %d", res.Code)
		}

	})

	t.Run("should parse query parameters and insert the result into the request context", func(t *testing.T) {
		email := "scooby@dooby.doo"
		next := func(w http.ResponseWriter, r *http.Request) {
			input := r.Context().Value("Input").(*Thing)

			if input.Email != email {
				t.Errorf("Thing was not successfully parsed")
			}

			w.WriteHeader(http.StatusOK)
		}

		req := httptest.NewRequest("GET", fmt.Sprintf("/?email=%s", email), nil)
		res := httptest.NewRecorder()
		middleware := ValidateMiddleware(utils, &Thing{}, next)

		middleware(res, req)

		if res.Code != 200 {
			t.Errorf("expected 200, got %d", res.Code)
		}
	})

	t.Run("should prioritize request body over query parameters when parsing and validating", func(t *testing.T) {
		email := "scooby@dooby.doo"
		next := func(w http.ResponseWriter, r *http.Request) {
			input := r.Context().Value("Input").(*Thing)

			if input.Email != email {
				t.Errorf("Thing was not successfully parsed")
			}

			w.WriteHeader(http.StatusOK)
		}

		reader := strings.NewReader(fmt.Sprintf(`{"email": "%s"}`, email))
		req := httptest.NewRequest("POST", "/?email=invalid", reader)
		req.Header.Add("content-type", "application/json")
		res := httptest.NewRecorder()
		middleware := ValidateMiddleware(utils, &Thing{}, next)

		middleware(res, req)

		if res.Code != 200 {
			t.Errorf("expected 200, got %d", res.Code)
		}

	})

	t.Run("should pass the struct on with no validation if there are no validation tags present", func(t *testing.T) {
		email := "scooby@dooby.doo"
		next := func(w http.ResponseWriter, r *http.Request) {
			input := r.Context().Value("Input").(*ThingNoValidation)

			if input.Email != email {
				t.Errorf("Thing was not successfully parsed")
			}

			w.WriteHeader(http.StatusOK)
		}

		reader := strings.NewReader(fmt.Sprintf(`{"email": "%s"}`, email))
		req := httptest.NewRequest("POST", "/", reader)
		req.Header.Add("content-type", "application/json")
		res := httptest.NewRecorder()
		middleware := ValidateMiddleware(utils, &ThingNoValidation{}, next)

		middleware(res, req)

		if res.Code != 200 {
			t.Errorf("expected 200, got %d", res.Code)
		}

	})

}
