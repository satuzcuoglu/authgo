package main

import (
	"net/http"

	"gym-back/database"
	"gym-back/routes"

	"github.com/gorilla/handlers"
)

func main() {
	database.SetupDB()
	router := routes.InitRoutes()
	headersOk := handlers.AllowedHeaders([]string{"Content-Type"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	http.ListenAndServe(":7000", handlers.CORS(originsOk, headersOk, methodsOk)(router))
}
