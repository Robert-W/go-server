package user

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/robert-w/go-server/internal/response"
)

// Create a mock that returns a successful response
type mockUserService struct{}

func (m *mockUserService) list(ctx context.Context) (*[]User, *response.ErrorJsonV1) {
	id, _ := uuid.NewV7()
	users := []User{{Id: id, Email: "Scooby Doo", Created: time.Now()}}
	return &users, nil
}

func (m *mockUserService) create(ctx context.Context, input *UserPost) (*[]User, *response.ErrorJsonV1) {
	id, _ := uuid.NewV7()
	users := []User{{Id: id, Email: input.Users[0].Email, Created: time.Now()}}
	return &users, nil
}

func (m *mockUserService) get(ctx context.Context, id string) (*User, *response.ErrorJsonV1) {
	parsed, _ := uuid.Parse(id)
	return &User{Id: parsed, Email: "Scooby Doo"}, nil
}

func (m *mockUserService) update(ctx context.Context, id string, input *UserPut) (*User, *response.ErrorJsonV1) {
	parsed, _ := uuid.Parse(id)
	fmt.Printf("ID %s", id)
	return &User{Id: parsed, Email: input.Email}, nil
}

func (m *mockUserService) delete(ctx context.Context, id string) (*User, *response.ErrorJsonV1) {
	parsed, _ := uuid.Parse(id)
	return &User{Id: parsed, Email: "Scooby Doo"}, nil
}

// Create a mock that returns a versioned error
type mockUserServiceErr struct{}

func (m *mockUserServiceErr) list(ctx context.Context) (*[]User, *response.ErrorJsonV1) {
	return nil, &response.ErrorJsonV1{Message: "Scooby Dooby Doo", StatusCode: 500, Original: errors.New("Mystery Inc")}
}

func (m *mockUserServiceErr) create(ctx context.Context, input *UserPost) (*[]User, *response.ErrorJsonV1) {
	return nil, &response.ErrorJsonV1{Message: "Scooby Dooby Doo", StatusCode: 500, Original: errors.New("Mystery Inc")}
}

func (m *mockUserServiceErr) get(ctx context.Context, id string) (*User, *response.ErrorJsonV1) {
	return nil, &response.ErrorJsonV1{Message: "Scooby Dooby Doo", StatusCode: 404, Original: errors.New("Mystery Inc")}
}

func (m *mockUserServiceErr) update(ctx context.Context, id string, input *UserPut) (*User, *response.ErrorJsonV1) {
	return nil, &response.ErrorJsonV1{Message: "Scooby Dooby Doo", StatusCode: 404, Original: errors.New("Mystery Inc")}
}

func (m *mockUserServiceErr) delete(ctx context.Context, id string) (*User, *response.ErrorJsonV1) {
	return nil, &response.ErrorJsonV1{Message: "Scooby Dooby Doo", StatusCode: 404, Original: errors.New("Mystery Inc")}
}

// Types for parsing responses
type mockResultUserList struct {
	Result []User `json:"result"`
}

type mockResultUser struct {
	Result User `json:"result"`
}

type mockResultV1Error struct {
	Error response.ErrorJsonV1 `json:"error"`
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
		email := "Scooby Doo"

		users := &UserPost{Users: []UserInput{{Email: email}}}
		ctx := context.WithValue(context.Background(), "Input", users)

		testHandler.create(res, req.WithContext(ctx))

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

		if result.Result[0].Email != email {
			t.Error("Result does not have the correct Name")
		}
	})

	t.Run("should return a v1Error if the underlying service returns an error", func(t *testing.T) {
		req := httptest.NewRequest("POST", "http://0.0.0.0:3000/api/v1/users", nil)
		res := httptest.NewRecorder()

		users := &UserPost{}
		ctx := context.WithValue(context.Background(), "Input", users)

		testHandlerErr.create(res, req.WithContext(ctx))

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
		id := uuid.New()
		url := fmt.Sprintf("http://0.0.0.0:3000/api/v1/users/%s", id.String())
		req := httptest.NewRequest("GET", url, nil)
		req = mux.SetURLVars(req, map[string]string{"id": id.String()})
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

		if result.Result.Id != id {
			t.Errorf("Expected %s, got %s", id.String(), result.Result.Id.String())
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
		id := uuid.New()
		url := fmt.Sprintf("http://0.0.0.0:3000/api/v1/users/%s", id.String())
		req := httptest.NewRequest("PUT", url, nil)
		res := httptest.NewRecorder()

		email := "Scooby Doo"
		user := &UserPut{Email: email}
		ctx := context.WithValue(context.Background(), "Input", user)
		req = req.WithContext(ctx)
		req = mux.SetURLVars(req, map[string]string{"id": id.String()})

		testHandler.update(res, req)

		if res.Code != 200 {
			t.Error("update should return a 200")
		}

		var result mockResultUser
		err := json.Unmarshal(res.Body.Bytes(), &result)

		if err != nil {
			t.Errorf("Unable to decode response: %v", err)
		}

		if result.Result.Email != email {
			t.Errorf("Expected email to be %s, got %s", email, result.Result.Email)
		}

		if result.Result.Id != id {
			t.Errorf("Expected id to be %s, got %s", id.String(), result.Result.Id.String())
		}
	})

	t.Run("should return a v1Error if the underlying service returns an error", func(t *testing.T) {
		req := httptest.NewRequest("PUT", "http://0.0.0.0:3000/api/v1/users/111", nil)
		res := httptest.NewRecorder()

		email := "Scooby Doo"
		user := &UserPut{Email: email}
		ctx := context.WithValue(context.Background(), "Input", user)

		testHandlerErr.update(res, req.WithContext(ctx))

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
		id := uuid.New()
		url := fmt.Sprintf("http://0.0.0.0:3000/api/v1/users/%s", id.String())
		req := httptest.NewRequest("DELETE", url, nil)
		req = mux.SetURLVars(req, map[string]string{"id": id.String()})
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
