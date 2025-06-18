package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sreekar2307/khata/model/table"
)

func (m middleware) UserAuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		var (
			bearerToken = context.GetHeader("Authorization")
			userService = m.userService
			user        table.User
		)
		if len(bearerToken) < 7 {
			context.AbortWithStatusJSON(401, gin.H{"error": "Authorization header is missing"})
			return
		}
		user, err := userService.ValidateAuthToken(context, bearerToken[7:])
		if err != nil {
			context.AbortWithStatusJSON(401, gin.H{"error": "Invalid or expired token"})
			return
		}
		context.Set("user", user)
	}
}
