package test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"testing"
)

// This file is for performing integration tests on the v1 user endpoints
type UserResponseMany struct {
	Status string `json:"status"`
	Result []User `json:"result"`
}

type UserResponseSingle struct {
	Status string `json:"status"`
	Result User   `json:"result"`
}

type User struct {
	Id    string `json:"id"`
	Email string `json:"email"`
}

func TestUserEndpoints(t *testing.T) {
	var err error
	var response *http.Response
	var request *http.Request
	var url string

	client := http.Client{}

	// Create a user
	reader := strings.NewReader(`{"users":[{"email":"scooby@dooby.doo"}]}`)
	response, err = http.Post("http://0.0.0.0:3000/api/v1/users", "application/json", reader)
	if err != nil {
		t.Errorf("Failed to create users: %v", err)
	}

	var result UserResponseMany
	json.NewDecoder(response.Body).Decode(&result)
	response.Body.Close()

	if result.Status != "ok" {
		t.Error("We did not get the correct type of result")
	}

	if result.Result[0].Email != "scooby@dooby.doo" {
		t.Errorf("Did not receive the expected email address, received %s", result.Result[0].Email)
	}

	userid := result.Result[0].Id

	// Update that user
	reader = strings.NewReader(`{"email":"scrappy@dappy.doo"}`)
	url = fmt.Sprintf("http://0.0.0.0:3000/api/v1/users/%s", userid)
	request, err = http.NewRequest(http.MethodPut, url, reader)
	if err != nil {
		t.Errorf("Failed to create put request: %v", err)
	}

	response, err = client.Do(request)
	if err != nil {
		t.Errorf("Failed to do put request: %v", err)
	}

	var singleResult UserResponseSingle
	json.NewDecoder(response.Body).Decode(&singleResult)
	response.Body.Close()

	if singleResult.Status != "ok" {
		t.Error("Failed to update user")
	}

	if singleResult.Result.Email != "scrappy@dappy.doo" {
		t.Errorf("Something in the update operation went wrong: %s", singleResult.Result.Email)
	}

	// Get the user by id
	url = fmt.Sprintf("http://0.0.0.0:3000/api/v1/users/%s", userid)
	response, err = http.Get(url)
	if err != nil {
		t.Errorf("Failed to get user: %v", err)
	}

	var getResult UserResponseSingle
	json.NewDecoder(response.Body).Decode(&getResult)
	response.Body.Close()

	if getResult.Status != "ok" {
		t.Error("We did not get the correct type of result")
	}

	if getResult.Result.Email != "scrappy@dappy.doo" {
		t.Errorf("Something in the get operation went wrong: %s", getResult.Result.Email)
	}

	// List all users
	response, err = http.Get("http://0.0.0.0:3000/api/v1/users")
	if err != nil {
		t.Errorf("Failed to list users: %v", err)
	}

	var listResult UserResponseMany
	json.NewDecoder(response.Body).Decode(&listResult)
	response.Body.Close()

	if listResult.Status != "ok" {
		t.Error("We did not get the correct type of result")
	}

	if listResult.Result[0].Email != "scrappy@dappy.doo" {
		t.Errorf("Did not receive the expected email address from listed user, received %s", result.Result[0].Email)
	}

	// Delete a user
	url = fmt.Sprintf("http://0.0.0.0:3000/api/v1/users/%s", userid)
	request, err = http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		t.Errorf("Failed to create delete request: %v", err)
	}

	response, err = client.Do(request)
	if err != nil {
		t.Errorf("Failed to delete user: %v", err)
	}

	var deleteResult UserResponseSingle
	json.NewDecoder(response.Body).Decode(&deleteResult)
	response.Body.Close()

	if deleteResult.Status != "ok" {
		t.Error("We did not get the correct type of result")
	}

	if deleteResult.Result.Email != "scrappy@dappy.doo" {
		t.Errorf("Something in the delete operation went wrong: %s", getResult.Result.Email)
	}
}
