package user

import (
	"context"
	"time"

	"github.com/robert-w/go-server/internal/monitoring"
	v1 "github.com/robert-w/go-server/internal/routes/v1"
	"go.opentelemetry.io/otel/codes"
)

type userService struct{}

func (s *userService) list(ctx context.Context) (*[]user, *v1.Error) {
	span := monitoring.CreateDBSpan(ctx, "UserService.list")
	defer span.End()

	users := []user{
		{
			Id:        "111",
			Value:     "First User",
			Timestamp: time.Now(),
		},
		{
			Id:        "222",
			Value:     "Second User",
			Timestamp: time.Now(),
		},
		{
			Id:        "333",
			Value:     "Third User",
			Timestamp: time.Now(),
		},
	}

	span.SetStatus(codes.Ok, "Ok")
	return &users, nil
}

func (s *userService) create(ctx context.Context) (*[]user, *v1.Error) {
	span := monitoring.CreateDBSpan(ctx, "UserService.create")
	defer span.End()

	users := []user{
		{
			Id:        "111",
			Value:     "New User",
			Timestamp: time.Now(),
		},
	}

	span.SetStatus(codes.Ok, "Ok")
	return &users, nil
}

func (s *userService) get(ctx context.Context) (*user, *v1.Error) {
	span := monitoring.CreateDBSpan(ctx, "UserService.get")
	defer span.End()

	user := user{
		Id:        "123",
		Value:     "User Read",
		Timestamp: time.Now(),
	}

	span.SetStatus(codes.Ok, "Ok")
	return &user, nil
}

func (s *userService) update(ctx context.Context) (*user, *v1.Error) {
	span := monitoring.CreateDBSpan(ctx, "UserService.update")
	defer span.End()

	user := user{
		Id:        "321",
		Value:     "User Update",
		Timestamp: time.Now(),
	}

	span.SetStatus(codes.Ok, "Ok")
	return &user, nil
}

func (s *userService) delete(ctx context.Context) (*user, *v1.Error) {
	span := monitoring.CreateDBSpan(ctx, "UserService.delete")
	defer span.End()

	user := user{
		Id: "321",
	}

	span.SetStatus(codes.Ok, "Ok")
	return &user, nil
}
