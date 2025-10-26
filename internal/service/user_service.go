package service

import (
	"context"
	"go-backend-task/internal/db"
	"time"
)

type UserService struct {
	queries *db.Queries
}

func NewUserService(q *db.Queries) *UserService {
	return &UserService{queries: q}
}

// CreateUser inserts a new user
func (s *UserService) CreateUser(ctx context.Context, name string, dob string) (db.User, error) {
	parsedDOB, err := time.Parse("2006-01-02", dob)
	if err != nil {
		return db.User{}, err
	}

	params := db.CreateUserParams{
		Name: name,
		Dob:  parsedDOB,
	}

	return s.queries.CreateUser(ctx, params)
}

// GetUserByID fetches a user by ID
func (s *UserService) GetUserByID(ctx context.Context, id int32) (db.User, error) {
	return s.queries.GetUserByID(ctx, id)
}

// GetAllUsers fetches all users
func (s *UserService) GetAllUsers(ctx context.Context) ([]db.User, error) {
	return s.queries.GetAllUsers(ctx)
}

// UpdateUser updates a user's name and dob
func (s *UserService) UpdateUser(ctx context.Context, id int32, name string, dob string) (db.User, error) {
	parsedDOB, err := time.Parse("2006-01-02", dob)
	if err != nil {
		return db.User{}, err
	}

	params := db.UpdateUserParams{
		ID:   id,
		Name: name,
		Dob:  parsedDOB,
	}

	return s.queries.UpdateUser(ctx, params)
}

// DeleteUser deletes a user by ID
func (s *UserService) DeleteUser(ctx context.Context, id int32) error {
	return s.queries.DeleteUser(ctx, id)
}
