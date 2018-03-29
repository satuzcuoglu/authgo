package controllers

import (
	"encoding/json"
	"net/http"

	"gym-back/models"
	"gym-back/services"
)

// Register helps users to register into application
func Register(rw http.ResponseWriter, req *http.Request) {
	var user models.User

	decodeError := json.NewDecoder(req.Body).Decode(&user)
	if decodeError != nil {
		panic(decodeError)
	}

	services.CreateUser(&user)

	userJSON, encodeError := json.Marshal(struct {
		models.User
		Password string `json:"-"`
	}{
		User: user,
	})
	if encodeError != nil {
		panic(encodeError)
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.Write(userJSON)
}

// Login helps user to login into application
func Login(rw http.ResponseWriter, req *http.Request) {
	var credentials loginCredentials

	decodeError := json.NewDecoder(req.Body).Decode(&credentials)
	if decodeError != nil {
		panic(decodeError)
	}

	result := services.CreateToken(credentials.Username, credentials.Password)
	if result == nil {
		rw.WriteHeader(http.StatusForbidden)
	} else {
		tokenJSON, encodeError := json.Marshal(result)
		if encodeError != nil {
			panic(encodeError)
		}

		rw.Header().Set("Content-Type", "application/json")
		rw.Write(tokenJSON)
	}
}

type loginCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
