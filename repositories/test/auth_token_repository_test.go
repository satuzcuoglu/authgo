package test

import (
	"go-auth/models"
	"go-auth/repositories"
	"testing"
	"time"

	"github.com/icrowley/fake"
)

var token = "qwerty12345"
var userData *models.User

func createTestUser() *models.User {
	user := new(models.User)
	user.Username = fake.UserName()
	user.Firstname = fake.FirstName()
	user.Lastname = fake.LastName()
	user.Password = fake.SimplePassword()
	user.Email = fake.EmailAddress()
	repositories.CreateUser(user)
	userData = user
	return user
}

func TestCreateAuthToken(t *testing.T) {
	authToken := new(models.AuthToken)
	authToken.Token = token
	authToken.User = *createTestUser()
	authToken.CreatedAt = time.Now().Add(72 * time.Hour)
	repositories.CreateAuthToken(authToken)

	if &authToken.User == nil {
		t.Error("Token doesnt have user info ")
	}
}

func TestFindAuthTokenByToken(t *testing.T) {
	authToken := repositories.FindAuthTokenByToken(token)
	userInfo := &authToken.User
	if userInfo == nil {
		t.Error("Token doesnt have user info")
	}
}
