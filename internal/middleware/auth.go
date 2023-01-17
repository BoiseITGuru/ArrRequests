package middleware

import (
	"github.com/BoiseITGuru/ArrRequests/internal/auth"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.GetHeader("Authorization")
		if tokenString == "" {
			context.AbortWithStatusJSON(401, gin.H{"error": "request does not contain an access token"})
			return
		}

		claims, err := auth.ValidateToken(tokenString)
		if err != nil {
			context.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
			return
		}

		context.Set("user", claims.ID)

		context.Next()
	}
}
