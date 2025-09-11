package user

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/robert-w/go-server/internal/monitoring"
	"github.com/robert-w/go-server/internal/response"
)

type userService struct{}

func (s *userService) list(ctx context.Context) (*[]User, *response.ErrorJsonV1) {
	_, span := monitoring.CreateDBSpan(ctx, "UserService.list")
	defer span.End()

	one, _ := uuid.NewV7()
	two, _ := uuid.NewV7()
	three, _ := uuid.NewV7()
	users := []User{
		{
			Id:      one,
			Email:    "First User",
			Created: time.Now(),
		},
		{
			Id:      two,
			Email:    "Second User",
			Created: time.Now(),
		},
		{
			Id:      three,
			Email:    "Third User",
			Created: time.Now(),
		},
	}

	return &users, nil
}

func (s *userService) create(ctx context.Context) (*[]User, *response.ErrorJsonV1) {
	_, span := monitoring.CreateDBSpan(ctx, "UserService.create")
	defer span.End()

	id, _ := uuid.NewV7()
	users := []User{
		{
			Id:      id,
			Email:    "New User",
			Created: time.Now(),
		},
	}

	return &users, nil
}

func (s *userService) get(ctx context.Context) (*User, *response.ErrorJsonV1) {
	_, span := monitoring.CreateDBSpan(ctx, "UserService.get")
	defer span.End()

	id, _ := uuid.NewV7()
	user := User{
		Id:      id,
		Email:    "User Read",
		Created: time.Now(),
	}

	return &user, nil
}

func (s *userService) update(ctx context.Context) (*User, *response.ErrorJsonV1) {
	_, span := monitoring.CreateDBSpan(ctx, "UserService.update")
	defer span.End()

	id, _ := uuid.NewV7()
	user := User{
		Id:      id,
		Email:    "User Update",
		Created: time.Now(),
	}

	return &user, nil
}

func (s *userService) delete(ctx context.Context) (*User, *response.ErrorJsonV1) {
	_, span := monitoring.CreateDBSpan(ctx, "UserService.delete")
	defer span.End()

	id, _ := uuid.NewV7()
	user := User{
		Id: id,
	}

	return &user, nil
}
