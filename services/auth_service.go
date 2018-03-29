package services

import (
	"crypto/rand"
	"fmt"
	"gym-back/database"
	"gym-back/models"
	"gym-back/repositories"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// CreateUser creates an user with hashed password
func CreateUser(user *models.User) *models.User {
	log.Println(user.Username)
	log.Println(user.Password)

	user.Password = generateHashedPassword(user.Password)
	repositories.CreateUser(user)
	return user
}

// CreateToken creates an auth_token
func CreateToken(username string, password string) *models.AuthToken {
	db := database.GetConnection()
	user := repositories.FindUserByUsername(username)
	token := new(models.AuthToken)

	log.Println(username)
	log.Println(password)

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	log.Println(err)
	if err != nil {
		log.Println("Error")
		return nil
	}

	token.Token = generateTokenString()
	token.User = *user
	token.ExpiresAt = time.Now().Add(72 * time.Hour)
	db.Create(&token)
	return token
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
