package repository

import (
	"context"

	"github.com/Luiggy102/go-rest-ws/models"
)

// two methods User
// insert
// get(by ID)
type UserRepo interface {
	InsertUser(ctx context.Context, user *models.User) error
	GetUserByID(ctx context.Context, id int64) (*models.User, error)
}

var implementation UserRepo

// UserRepo is equal at any implementation of the interface
// different databases for example
func SetRepo(repository UserRepo) { // postgres implementation, mongodb implementation
	// dependency injection
	implementation = repository
}

func InsertUser(ctx context.Context, user *models.User) error {
	return implementation.InsertUser(ctx, user) // implementation
}

func GetUserByID(ctx context.Context, id int64) (*models.User, error) {
	return implementation.GetUserByID(ctx, id) // implementation
}
