package main

import (
	"go-auth/database"
	"go-auth/routes"
)

func main() {
	database.SetupDB()
	routes.InitRoutes()
}
