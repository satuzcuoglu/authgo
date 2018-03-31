package test

import (
	"testing"

	"github.com/icrowley/fake"

	"gym-back/models"
	"gym-back/repositories"
)

var (
	userID    uint
	userName  = fake.UserName()
	email     = fake.EmailAddress()
	firstName = fake.FirstName()
	lastName  = fake.LastName()
	password  = fake.SimplePassword()
)

func TestCreateUser(t *testing.T) {
	user := new(models.User)
	user.Username = userName
	user.Firstname = firstName
	user.Lastname = lastName
	user.Password = password
	user.Email = email
	repositories.CreateUser(user)
	userID = user.ID
}

func TestFindUserByID(t *testing.T) {
	user := repositories.FindUserByID(userID)
	if user.ID != userID {
		t.Error("user ID's doesnt match , got ", user.ID)
	}
}

func TestFindUserByUsername(t *testing.T) {
	user := repositories.FindUserByUsername(userName)
	if user == nil {
		t.Error("User cannot found via username", userName)
	}
}

func TestFindUserByEmail(t *testing.T) {
	user := repositories.FindUserByEmail(email)
	if user == nil {
		t.Error("User cannot found via email", email)
	}
}

func TestDeleteUser(t *testing.T) {
	user := repositories.FindUserByID(1)
	repositories.DeleteUser(user)
	if user.DeletedAt != nil {
		t.Error("User Cannot Deleted")
	}
}
