package routes

import (
	"authgo/config"
	"authgo/controllers"
	"strconv"

	"github.com/gin-gonic/gin"
)

// InitRoutes initializes all routes for app
func InitRoutes() {
	conf := config.GetConfig()
	router := gin.Default()
	router.Use(CORSMiddleware())

	auth := router.Group("/auth")
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)
		auth.GET("/logout", controllers.Logout)
		auth.Use(AuthMiddleware()).GET("/me", controllers.User)
	}

	router.Run(":" + strconv.Itoa(conf.Port))
}
