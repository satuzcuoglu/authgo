package routes

import (
	"gym-back/repositories"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware checks accesibility on protected routes
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Token")
		token := repositories.FindAuthTokenByToken(tokenString)
		if token.UserID == 0 {
			c.AbortWithStatus(401)
			return
		}
		c.Set("user", token.User)
		c.Next()
	}
}
