package user

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/noir-cats/go-sample-server/internal/monitoring"
	"github.com/noir-cats/go-sample-server/internal/response"
	"go.opentelemetry.io/otel/codes"
)

type userService struct {
	pool *pgxpool.Pool
}

func (s *userService) list(ctx context.Context) (*[]User, *response.ErrorJsonV1) {
	ctx, span := monitoring.CreateDBSpan(ctx, "UserService.list")
	defer span.End()

	rows, _ := s.pool.Query(ctx, "SELECT * FROM users")
	users, err := pgx.CollectRows(rows, pgx.RowToStructByName[User])
	if err != nil {
		dberr := &response.ErrorJsonV1{
			Message:    "Error fetching all users from pool",
			Original:   err,
			StatusCode: http.StatusInternalServerError,
		}

		span.RecordError(dberr.Original)
		span.SetStatus(codes.Error, dberr.Original.Error())
		return nil, dberr
	}

	return &users, nil
}

func (s *userService) create(ctx context.Context, userInput *UserPost) (*[]User, *response.ErrorJsonV1) {
	ctx, span := monitoring.CreateDBSpan(ctx, "UserService.create")
	defer span.End()

	// Create a list of users to insert
	var users []User = make([]User, len(userInput.Users))
	for idx, item := range userInput.Users {
		id, _ := uuid.NewV7()
		now := time.Now()
		users[idx] = User{
			Id:          id,
			Email:       item.Email,
			Created:     now,
			LastUpdated: now,
		}
	}

	// This is the recommended way to do a multi-insert, it is generally more
	// efficient if the insert includes more than 5 users. If < 5, the performance
	// penalty is likely not enough to be concerned with (but can verify later)
	_, err := s.pool.CopyFrom(
		ctx,
		pgx.Identifier{"users"},
		[]string{"id", "email", "created", "last_updated"},
		pgx.CopyFromSlice(len(users), func(i int) ([]any, error) {
			return []any{users[i].Id, users[i].Email, users[i].Created, users[i].LastUpdated}, nil
		}),
	)

	if err != nil {
		dberr := &response.ErrorJsonV1{
			Message:    "Error inserting users",
			Original:   err,
			StatusCode: http.StatusInternalServerError,
		}

		span.RecordError(dberr.Original)
		span.SetStatus(codes.Error, dberr.Original.Error())
		return nil, dberr
	}

	return &users, nil
}

func (s *userService) get(ctx context.Context, id string) (*User, *response.ErrorJsonV1) {
	_, span := monitoring.CreateDBSpan(ctx, "UserService.get")
	defer span.End()

	rows, _ := s.pool.Query(ctx, "SELECT * FROM users WHERE id = $1", id)
	user, err := pgx.CollectExactlyOneRow(rows, pgx.RowToAddrOfStructByName[User])
	if err != nil {
		dberr := &response.ErrorJsonV1{
			Message:    fmt.Sprintf("Error fetching user with id = %s", id),
			Original:   err,
			StatusCode: http.StatusInternalServerError,
		}

		span.RecordError(dberr.Original)
		span.SetStatus(codes.Error, dberr.Original.Error())
		return nil, dberr
	}

	return user, nil
}

func (s *userService) update(ctx context.Context, id string, input *UserPut) (*User, *response.ErrorJsonV1) {
	_, span := monitoring.CreateDBSpan(ctx, "UserService.update")
	defer span.End()

	now := time.Now()
	rows, _ := s.pool.Query(ctx, "UPDATE users SET email = $1, last_updated = $2 WHERE id = $3 RETURNING *", input.Email, now, id)
	user, err := pgx.CollectExactlyOneRow(rows, pgx.RowToAddrOfStructByName[User])
	if err != nil {
		dberr := &response.ErrorJsonV1{
			Message:    fmt.Sprintf("Error updating user with id = %s", id),
			Original:   err,
			StatusCode: http.StatusInternalServerError,
		}

		span.RecordError(dberr.Original)
		span.SetStatus(codes.Error, dberr.Original.Error())
		return nil, dberr
	}

	return user, nil
}

func (s *userService) delete(ctx context.Context, id string) (*User, *response.ErrorJsonV1) {
	_, span := monitoring.CreateDBSpan(ctx, "UserService.delete")
	defer span.End()

	rows, err := s.pool.Query(ctx, "DELETE FROM users WHERE id = $1 RETURNING *", id)
	user, err := pgx.CollectExactlyOneRow(rows, pgx.RowToAddrOfStructByName[User])
	if err != nil {
		dberr := &response.ErrorJsonV1{
			Message:    fmt.Sprintf("Error deleting user with id = %s", id),
			Original:   err,
			StatusCode: http.StatusInternalServerError,
		}

		span.RecordError(dberr.Original)
		span.SetStatus(codes.Error, dberr.Original.Error())
		return nil, dberr
	}

	return user, nil
}
