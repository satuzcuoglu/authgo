package routes

import (
	"gym-back/config"
	"gym-back/controllers"
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
	}

	apiRoutes := router.Group("/api")
	apiRoutes.Use(AuthMiddleware())
	{
		apiRoutes.GET("/user", controllers.User)
	}

	router.Run(":" + strconv.Itoa(conf.Port))
}
