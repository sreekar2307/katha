package user

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/sreekar2307/katha/errors"
	"github.com/sreekar2307/katha/model/table"
	"github.com/sreekar2307/katha/pkg/jwt"
	"github.com/sreekar2307/katha/repository"
	"github.com/sreekar2307/katha/service"
	"gorm.io/gorm"
	"time"
)

type userServ struct {
	primaryDB  *gorm.DB
	repository repository.Repository
	jwtImpl    jwt.JWT
}

func NewUserService(
	primaryDB *gorm.DB,
	repository repository.Repository,
	jwtImpl jwt.JWT,
) service.User {
	return userServ{
		primaryDB:  primaryDB,
		repository: repository,
		jwtImpl:    jwtImpl,
	}
}

func (u userServ) AddUser(ctx context.Context, user table.User) (table.User, error) {
	if len(user.Password) == 0 {
		return table.User{}, fmt.Errorf("password cannot be empty")
	}
	user.PasswordHash = hashPassword(user.Password)
	if err := u.repository.CreateUser(ctx, u.primaryDB, &user); err != nil {
		return table.User{}, fmt.Errorf("failed to create user: %w", err)
	}
	return user, nil
}

func (u userServ) Authenticate(ctx context.Context, email, password string) (table.User, string, error) {
	passwordHash := hashPassword(password)
	user, err := u.repository.GetUserByEmail(ctx, u.primaryDB, email)
	if err != nil {
		return table.User{}, "", fmt.Errorf("failed to get user by email: %w", err)
	}
	if user.PasswordHash != passwordHash {
		return table.User{}, "", errors.ErrInvalidUserCredentials
	}
	token, err := u.jwtImpl.Token(ctx, map[string]any{
		"email": user.Email,
		"exp":   time.Now().Add(24 * time.Hour).Unix(),
	})
	if err != nil {
		return table.User{}, "", fmt.Errorf("failed to generate auth token: %w", err)
	}
	return user, token, nil
}

func (u userServ) ValidateAuthToken(ctx context.Context, token string) (table.User, error) {
	claims, err := u.jwtImpl.Validate(ctx, token)
	if err != nil {
		return table.User{}, fmt.Errorf("failed to validate auth token: %w", err)
	}
	email, ok := claims["email"].(string)
	if !ok {
		return table.User{}, fmt.Errorf("invalid token claims: missing email")
	}
	user, err := u.repository.GetUserByEmail(ctx, u.primaryDB, email)
	if err != nil {
		return table.User{}, fmt.Errorf("failed to get user by email: %w", err)
	}
	return user, nil
}

func hashPassword(password string) string {
	return hex.EncodeToString(sha256.New().Sum([]byte(password)))
}
