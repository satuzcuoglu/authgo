package main

import (
	"github.com/satuzcuoglu/authgo/database"
	"github.com/satuzcuoglu/authgo/routes"
)

func main() {
	database.SetupDB()
	routes.InitRoutes()
}
