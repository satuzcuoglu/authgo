package controllers

import (
	"net/http"

	"gym-back/models"
	"gym-back/services"

	"github.com/gin-gonic/gin"
)

// Register helps users to register into application
func Register(c *gin.Context) {
	var user *models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		user, registerError := services.CreateUser(user)
		if registerError != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": registerError.Error()})
		} else {
			c.JSON(http.StatusOK, user)
		}
	}
}

// Login helps user to login into application
func Login(c *gin.Context) {
	var credentials loginCredentials

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		token, err := services.CreateToken(credentials.Username, credentials.Password)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, token)
		}
	}
}

type loginCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
