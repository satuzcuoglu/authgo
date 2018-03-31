package main

import (
	"gym-back/database"
	"gym-back/routes"
)

func main() {
	database.SetupDB()
	routes.InitRoutes()
}
