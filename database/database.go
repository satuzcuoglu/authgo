package database

import (
	"github.com/jinzhu/gorm"

	// Postgresql Dialect
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

func createConnection() *gorm.DB {
	db, err := gorm.Open("postgres", "host=localhost port=5432 dbname=gymsub_dev user=gymsub password=gymsub sslmode=disable")
	if err != nil {
		panic(err)
	}
	return db
}

// GetConnection returns and database connection
func GetConnection() *gorm.DB {
	if db != nil {
		return db
	}
	return createConnection()
}

// SetupDB Create Tables And Generates Migration
func SetupDB() {
	db = createConnection()
	db.LogMode(true)
	autoMigrateDatabases()
}
