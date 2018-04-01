package repositories

import (
	"gym-back/database"
	"gym-back/models"
)

// CreateUser creates an user and returns 201
func CreateUser(user *models.User) *models.User {
	db := database.GetConnection()
	db.Create(&user)
	return user
}

// FindUserByID finds user by id and returns it
func FindUserByID(id uint) *models.User {
	db := database.GetConnection()
	user := new(models.User)
	db.Preload("Authorities").Where("id = ?", id).First(&user)
	return user
}

// DeleteUser deletes an user
func DeleteUser(user *models.User) *models.User {
	db := database.GetConnection()
	if user.Email != "" && user.Username != "" {
		db.Delete(&user)
	}
	return user
}

// FindUserByUsername finds user by username and returns it
func FindUserByUsername(username string) *models.User {
	db := database.GetConnection()
	user := new(models.User)
	db.Unscoped().Preload("Authorities").Where("username = ?", username).First(&user)
	return user
}

// FindUserByEmail finds user by email and returns it
func FindUserByEmail(email string) *models.User {
	db := database.GetConnection()
	user := new(models.User)
	db.Unscoped().Preload("Authorities").Where("email = ?", email).First(&user)
	return user
}
