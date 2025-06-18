package service

import (
	"context"
	"github.com/sreekar2307/katha/model/table"
)

type User interface {
	AddUser(ctx context.Context, user table.User) (table.User, error)
	Authenticate(ctx context.Context, email, password string) (table.User, string, error)
	ValidateAuthToken(ctx context.Context, token string) (table.User, error)
}
