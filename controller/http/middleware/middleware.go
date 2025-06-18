package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sreekar2307/khata/service"
)

type middleware struct {
	userService service.User
}

type Middleware interface {
	// UserAuthMiddleware returns a gin.HandlerFunc that validates the user's bearer token.
	UserAuthMiddleware() gin.HandlerFunc
}

func NewMiddleware(userService service.User) Middleware {
	return &middleware{
		userService: userService,
	}
}
