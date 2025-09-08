package user

import (
	"context"
	"encoding/json"
	"errors"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/google/uuid"
	v1 "github.com/robert-w/go-server/internal/routes/v1"
)

// Create a mock that returns a successful response
type mockUserService struct{}

func (m *mockUserService) list(ctx context.Context) (*[]user, *v1.Error) {
	id, _ := uuid.NewV7()
	users := []user{{Id: id, Email: "Scooby Doo", Created: time.Now()}}
	return &users, nil
}

func (m *mockUserService) create(ctx context.Context) (*[]user, *v1.Error) {
	id, _ := uuid.NewV7()
	users := []user{{Id: id, Email: "Scooby Doo", Created: time.Now()}}
	return &users, nil
}

func (m *mockUserService) get(ctx context.Context) (*user, *v1.Error) {
	id, _ := uuid.NewV7()
	return &user{Id: id, Email: "Scooby Doo"}, nil
}

func (m *mockUserService) update(ctx context.Context) (*user, *v1.Error) {
	id, _ := uuid.NewV7()
	return &user{Id: id, Email: "Scooby Doo"}, nil
}

func (m *mockUserService) delete(ctx context.Context) (*user, *v1.Error) {
	id, _ := uuid.NewV7()
	return &user{Id: id, Email: "Scooby Doo"}, nil
}

// Create a mock that returns a versioned error
type mockUserServiceErr struct{}

func (m *mockUserServiceErr) list(ctx context.Context) (*[]user, *v1.Error) {
	return nil, &v1.Error{Message: "Scooby Dooby Doo", StatusCode: 500, Original: errors.New("Mystery Inc")}
}

func (m *mockUserServiceErr) create(ctx context.Context) (*[]user, *v1.Error) {
	return nil, &v1.Error{Message: "Scooby Dooby Doo", StatusCode: 500, Original: errors.New("Mystery Inc")}
}

func (m *mockUserServiceErr) get(ctx context.Context) (*user, *v1.Error) {
	return nil, &v1.Error{Message: "Scooby Dooby Doo", StatusCode: 404, Original: errors.New("Mystery Inc")}
}

func (m *mockUserServiceErr) update(ctx context.Context) (*user, *v1.Error) {
	return nil, &v1.Error{Message: "Scooby Dooby Doo", StatusCode: 404, Original: errors.New("Mystery Inc")}
}

func (m *mockUserServiceErr) delete(ctx context.Context) (*user, *v1.Error) {
	return nil, &v1.Error{Message: "Scooby Dooby Doo", StatusCode: 404, Original: errors.New("Mystery Inc")}
}

// Types for parsing responses
type mockResultUserList struct {
	Result []user `json:"result"`
}

type mockResultUser struct {
	Result user `json:"result"`
}

type mockResultV1Error struct {
	Error v1.Error `json:"error"`
}

func TestListUsers(t *testing.T) {
	testHandler := handler{service: &mockUserService{}}
	testHandlerErr := handler{service: &mockUserServiceErr{}}

	t.Run("should return users in the format of a v1Response", func(t *testing.T) {
		req := httptest.NewRequest("GET", "http://0.0.0.0:3000/api/v1/users", nil)
		res := httptest.NewRecorder()

		testHandler.list(res, req)

		if res.Code != 200 {
			t.Error("list should return a 200")
		}

		var result mockResultUserList
		err := json.Unmarshal(res.Body.Bytes(), &result)

		if err != nil {
			t.Errorf("Unable to decode response: %v", err)
		}

		if len(result.Result) != 1 {
			t.Error("Result is not the correct length")
		}

		if result.Result[0].Email != "Scooby Doo" {
			t.Error("Result does not have the correct Name")
		}
	})

	t.Run("should return a v1Error if the underlying service returns an error", func(t *testing.T) {
		req := httptest.NewRequest("GET", "http://0.0.0.0:3000/api/v1/users", nil)
		res := httptest.NewRecorder()

		testHandlerErr.list(res, req)

		if res.Code != 500 {
			t.Errorf("list should return a 500, got %d", res.Code)
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

func TestCreateUsers(t *testing.T) {
	testHandler := handler{service: &mockUserService{}}
	testHandlerErr := handler{service: &mockUserServiceErr{}}

	t.Run("should return the created users in the format of a v1Response", func(t *testing.T) {
		req := httptest.NewRequest("POST", "http://0.0.0.0:3000/api/v1/users", nil)
		res := httptest.NewRecorder()

		testHandler.create(res, req)

		if res.Code != 200 {
			t.Error("create should return a 200")
		}

		var result mockResultUserList
		err := json.Unmarshal(res.Body.Bytes(), &result)

		if err != nil {
			t.Errorf("Unable to decode response: %v", err)
		}

		if len(result.Result) != 1 {
			t.Error("Result is not the correct length")
		}

		if result.Result[0].Email != "Scooby Doo" {
			t.Error("Result does not have the correct Name")
		}
	})

	t.Run("should return a v1Error if the underlying service returns an error", func(t *testing.T) {
		req := httptest.NewRequest("POST", "http://0.0.0.0:3000/api/v1/users", nil)
		res := httptest.NewRecorder()

		testHandlerErr.create(res, req)

		if res.Code != 500 {
			t.Errorf("create should return a 500, got %d", res.Code)
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

func TestGetUser(t *testing.T) {
	testHandler := handler{service: &mockUserService{}}
	testHandlerErr := handler{service: &mockUserServiceErr{}}

	t.Run("should return the user in the format of a v1Response", func(t *testing.T) {
		req := httptest.NewRequest("GET", "http://0.0.0.0:3000/api/v1/users/111", nil)
		res := httptest.NewRecorder()

		testHandler.get(res, req)

		if res.Code != 200 {
			t.Error("get should return a 200")
		}

		var result mockResultUser
		err := json.Unmarshal(res.Body.Bytes(), &result)

		if err != nil {
			t.Errorf("Unable to decode response: %v", err)
		}

		if result.Result.Email != "Scooby Doo" {
			t.Error("Result does not have the correct Name")
		}
	})

	t.Run("should return a v1Error if the underlying service returns an error", func(t *testing.T) {
		req := httptest.NewRequest("GET", "http://0.0.0.0:3000/api/v1/users/111", nil)
		res := httptest.NewRecorder()

		testHandlerErr.get(res, req)

		if res.Code != 404 {
			t.Errorf("get should return a 404, got %d", res.Code)
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

func TestUpdateUser(t *testing.T) {
	testHandler := handler{service: &mockUserService{}}
	testHandlerErr := handler{service: &mockUserServiceErr{}}

	t.Run("should return the updated user in the format of a v1Response", func(t *testing.T) {
		req := httptest.NewRequest("PUT", "http://0.0.0.0:3000/api/v1/users/111", nil)
		res := httptest.NewRecorder()

		testHandler.update(res, req)

		if res.Code != 200 {
			t.Error("update should return a 200")
		}

		var result mockResultUser
		err := json.Unmarshal(res.Body.Bytes(), &result)

		if err != nil {
			t.Errorf("Unable to decode response: %v", err)
		}

		if result.Result.Email != "Scooby Doo" {
			t.Error("Result does not have the correct Name")
		}
	})

	t.Run("should return a v1Error if the underlying service returns an error", func(t *testing.T) {
		req := httptest.NewRequest("PUT", "http://0.0.0.0:3000/api/v1/users/111", nil)
		res := httptest.NewRecorder()

		testHandlerErr.update(res, req)

		if res.Code != 404 {
			t.Errorf("update should return a 404, got %d", res.Code)
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

func TestDeleteUser(t *testing.T) {
	testHandler := handler{service: &mockUserService{}}
	testHandlerErr := handler{service: &mockUserServiceErr{}}

	t.Run("should return the id of the deleted user in the format of a v1Response", func(t *testing.T) {
		req := httptest.NewRequest("DELETE", "http://0.0.0.0:3000/api/v1/users/111", nil)
		res := httptest.NewRecorder()

		testHandler.delete(res, req)

		if res.Code != 200 {
			t.Error("delete should return a 200")
		}

		var result mockResultUser
		err := json.Unmarshal(res.Body.Bytes(), &result)

		if err != nil {
			t.Errorf("Unable to decode response: %v", err)
		}

		if result.Result.Email != "Scooby Doo" {
			t.Error("Result does not have the correct Name")
		}
	})

	t.Run("should return a v1Error if the underlying service returns an error", func(t *testing.T) {
		req := httptest.NewRequest("DELETE", "http://0.0.0.0:3000/api/v1/users/111", nil)
		res := httptest.NewRecorder()

		testHandlerErr.delete(res, req)

		if res.Code != 404 {
			t.Errorf("delete should return a 404, got %d", res.Code)
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
