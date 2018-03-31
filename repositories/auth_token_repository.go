package repositories

import (
	"gym-back/database"
	"gym-back/models"
	"time"
)

// CreateAuthToken generates and token for an user
func CreateAuthToken(token *models.AuthToken) *models.AuthToken {
	db := database.GetConnection()
	db.Create(&token)
	return token
}

// FindUserByToken finds User by token and returns it
func FindUserByToken(token string) *models.AuthToken {
	db := database.GetConnection()
	authToken := new(models.AuthToken)
	db.Preload("User").Where("token = ? and expires_at > ?", token, time.Now()).First(&authToken)
	return authToken
}
