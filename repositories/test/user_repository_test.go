package test

import (
	"testing"

	"gym-back/models"
	"gym-back/repositories"
)

func TestCreateUser(t *testing.T) {
	user := new(models.User)
	user.Username = "john"
	user.Firstname = "John"
	user.Lastname = "Doe"
	user.Password = "test123"
	user.Email = "john.doe@test.com"
	repositories.CreateUser(user)

	if user.ID != 1 {
		t.Error("Expected 1, got ", user.ID)
	}
}

func TestFindUserByID(t *testing.T) {
	user := repositories.FindUserByID(1)
	if user.ID != 1 {
		t.Error("Expected 1, got ", user.ID)
	}
}

func TestDeleteUser(t *testing.T) {
	user := repositories.FindUserByID(1)
	repositories.DeleteUser(user)
	if user.DeletedAt != nil {
		t.Error("User Cannot Deleted")
	}
}
