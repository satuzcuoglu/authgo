package repositories

import (
	"time"

	"github.com/satuzcuoglu/authgo/database"
	"github.com/satuzcuoglu/authgo/models"
)

// CreateAuthToken generates and token for an user
func CreateAuthToken(token *models.AuthToken) *models.AuthToken {
	db := database.GetConnection()
	db.Create(&token)
	return token
}

// DeleteAuthToken deletes generated token
func DeleteAuthToken(token string) {
	db := database.GetConnection()
	db.Where("token = ?", token).Delete(models.AuthToken{})
}

// FindAuthTokenByToken finds User by token and returns it
func FindAuthTokenByToken(token string) *models.AuthToken {
	db := database.GetConnection()
	authToken := new(models.AuthToken)
	db.Preload("User").Preload("User.Authorities").Where("token = ? and expires_at > ?", token, time.Now()).First(&authToken)
	return authToken
}
