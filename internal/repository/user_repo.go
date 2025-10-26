package repository

import (
	"context"
	"time"

	db "go-backend-task/db/sqlc"
)

// UserRepository handles user-related DB operations
type UserRepository struct {
	queries *db.Queries
}

// NewUserRepository creates a new UserRepository
func NewUserRepository(q *db.Queries) *UserRepository {
	return &UserRepository{
		queries: q,
	}
}

// CreateUser inserts a new user into the database
func (r *UserRepository) CreateUser(ctx context.Context, name string, dob time.Time) (db.User, error) {
	params := db.CreateUserParams{
		Name: name,
		Dob:  dob, // time.Time matches SQLC generated struct
	}

	return r.queries.CreateUser(ctx, params)
}

// GetUserByID retrieves a user by ID
func (r *UserRepository) GetUserByID(ctx context.Context, id int64) (db.User, error) {
	return r.queries.GetUserByID(ctx, id)
}
