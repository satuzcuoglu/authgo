package routes

import (
	"gym-back/controllers"

	"github.com/gorilla/mux"
)

// InitRoutes initializes all routes for app
func InitRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/auth/register", controllers.Register)
	router.HandleFunc("/auth/login", controllers.Login)
	return router
}
