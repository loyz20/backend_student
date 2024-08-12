package middleware

import (
	"backend_student/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.GetHeader("Authorization")
		if tokenStr == "" {
			utils.RespondError(c, http.StatusUnauthorized, "No token provided", nil)
			c.Abort()
			return
		}

		claims, err := utils.ParseToken(tokenStr)
		if err != nil {
			utils.RespondError(c, http.StatusUnauthorized, "Invalid token", err)
			c.Abort()
			return
		}

		// Save claims to context for later use
		username, ok := claims["username"].(string)
		if !ok {
			utils.RespondError(c, http.StatusUnauthorized, "Invalid token claims", nil)
			c.Abort()
			return
		}

		c.Set("username", username)

		c.Next()
	}
}
