package services

import (
	"crypto/rand"
	"errors"
	"fmt"
	"gym-back/database"
	"gym-back/models"
	"gym-back/repositories"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// CreateUser creates an user with hashed password
func CreateUser(user *models.User) (*models.User, error) {
	userByUsername := repositories.FindUserByUsername(user.Username)
	if userByUsername.Username != "" {
		return nil, errors.New("Username Already Taken")
	}

	userByEmail := repositories.FindUserByEmail(user.Email)
	if userByEmail.Email != "" {
		return nil, errors.New("Email Already in use")
	}

	user.Password = generateHashedPassword(user.Password)
	repositories.CreateUser(user)
	return user, nil
}

// CreateToken creates an auth_token
func CreateToken(username string, password string) (*models.AuthToken, error) {
	db := database.GetConnection()
	user := repositories.FindUserByUsername(username)
	token := new(models.AuthToken)

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, err
	}

	token.Token = generateTokenString()
	token.User = *user
	token.ExpiresAt = time.Now().Add(72 * time.Hour)
	db.Create(&token)
	return token, nil
}

func generateHashedPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hashedPassword)
}

func generateTokenString() string {
	b := make([]byte, 20)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
