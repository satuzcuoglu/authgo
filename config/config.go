package config

import (
	"encoding/json"
	"os"
)

// Config struct keeps config variables
type Config struct {
	Port       int
	DBHost     string
	DBPort     int
	DBName     string
	DBUsername string
	DBPassword string
}

// GetConfig return app configuration File
func GetConfig() Config {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}

	file, err := os.Open(os.Getenv("GOPATH") + "/src/gym-back/config/config." + env + ".json")
	if err != nil {
		panic(err)
	}

	config := new(Config)
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		panic(err)
	}
	return *config
}

func getWd() string {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return wd
}
