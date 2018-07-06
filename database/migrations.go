package database

import (
	"github.com/satuzcuoglu/authgo/models"
)

func autoMigrateDatabases() {
	userAuthority := new(models.Authority)
	userAuthority.Name = "ROLE_USER"
	db.AutoMigrate(models.User{}, models.AuthToken{}, models.Authority{})
	db.Create(userAuthority)
}
