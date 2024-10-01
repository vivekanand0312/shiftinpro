package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shiftinpro/utility"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": false, "message": "Authorization header required"})
			c.Abort()
			return
		}

		tokenString := authHeader[len("Bearer "):]
		claims, err := utility.ValidateJWT(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"status": false, "message": "Invalid token", "error": err.Error()})
			c.Abort()
			return
		}

		c.Set("phone", claims.Phone) // Or set a different claim as needed
		c.Next()
	}
}
