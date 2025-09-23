package user

import (
	"context"
	"testing"
)

// These tests are placeholders as the service is currently hardcoded to return
// a list of users, once we use a real database, this will need to change
func TestUserService(t *testing.T) {
	ctx := context.Background()
	service := userService{}

	// TODO: Provide mock for pool so that these tests dont fail
	t.Skip("Skipping service tests until I mock the pool")

	t.Run("should return a list of users", func(t *testing.T) {
		users, _ := service.list(ctx)

		if len(*users) != 3 {
			t.Errorf("Expected three users, got %d", len(*users))
		}
	})

	t.Run("should create a set of users", func(t *testing.T) {
		input := &UserPost{}
		users, _ := service.create(ctx, input)

		if len(*users) != 1 {
			t.Errorf("Expected one user, got %d", len(*users))
		}
	})

	t.Run("should read a single user", func(t *testing.T) {
		user, _ := service.get(ctx, "")

		if user == nil {
			t.Error("Expected user, got nil")
		}
	})

	t.Run("should update a single user", func(t *testing.T) {
		input := &UserPut{}
		user, _ := service.update(ctx, "", input)

		if user == nil {
			t.Error("Expected user, got nil")
		}
	})

	t.Run("should delete a single user", func(t *testing.T) {
		user, _ := service.delete(ctx, "")

		if user == nil {
			t.Error("Expected user, got nil")
		}
	})
}
