package service

import (
	"context"
	"github.com/sreekar2307/katha/model/table"
)

// User is an interface that defines methods for user management, including adding users.
type User interface {
	// AddUser adds a new user to the system and returns the created user.
	AddUser(ctx context.Context, user table.User) (table.User, error)
	// Authenticate checks the user's credentials and returns the user and an authentication token if successful.
	Authenticate(ctx context.Context, email, password string) (table.User, string, error)
	// ValidateAuthToken checks the validity of the provided authentication token and returns the user if valid.
	ValidateAuthToken(ctx context.Context, token string) (table.User, error)
}
