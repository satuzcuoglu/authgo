package repositories

import (
	"go-auth/database"

	"go-auth/models"
)

// FindAuthorityByName returns an authority
func FindAuthorityByName(name string) *models.Authority {
	db := database.GetConnection()
	authority := new(models.Authority)
	db.Where("name = ?", name).First(&authority)
	return authority
}
