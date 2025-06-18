package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sreekar2307/katha/service"
)

type middleware struct {
	userService service.User
}

type Middleware interface {
	UserAuthMiddleware() gin.HandlerFunc
}

func NewMiddleware(userService service.User) Middleware {
	return &middleware{
		userService: userService,
	}
}
