package repositories

import (
	"gym-back/database"

	"gym-back/models"
)

// FindAuthorityByName returns an authority
func FindAuthorityByName(name string) *models.Authority {
	db := database.GetConnection()
	authority := new(models.Authority)
	db.Where("name = ?", name).First(&authority)
	return authority
}
