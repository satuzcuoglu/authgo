package main

import (
	"authgo/database"
	"authgo/routes"
)

func main() {
	database.SetupDB()
	routes.InitRoutes()
}
