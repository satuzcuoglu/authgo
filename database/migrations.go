package database

import (
	"gym-back/models"
)

func autoMigrateDatabases() {
	// db.DropTableIfExists(models.User{}, models.AuthToken{})
	db.AutoMigrate(models.User{}, models.AuthToken{})
}
