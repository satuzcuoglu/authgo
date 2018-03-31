package routes

import (
	"gym-back/controllers"

	"github.com/gin-gonic/gin"
)

// InitRoutes initializes all routes for app
func InitRoutes() {
	router := gin.Default()
	router.Use(CORSMiddleware())
	router.POST("/auth/register", controllers.Register)
	router.POST("/auth/login", controllers.Login)
	router.Run(":7000")
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
