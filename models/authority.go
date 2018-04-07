package models

// Authority returns users authorities
type Authority struct {
	ID   uint   `json:"-" binding:"required"`
	Name string `gorm:"unique_index" json:"name" binding:"required"`
}
