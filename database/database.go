package database

import (
	"fmt"

	"github.com/satuzcuoglu/authgo/config"

	"github.com/jinzhu/gorm"

	// Postgresql Dialect
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

func createConnection() *gorm.DB {
	conf := config.GetConfig()
	connectionString :=
		fmt.Sprintf("host=%s port=%v dbname=%s user=%s password=%s sslmode=disable",
			conf.DBHost, conf.DBPort, conf.DBName, conf.DBUsername, conf.DBPassword)
	db, err := gorm.Open("postgres", connectionString)
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
