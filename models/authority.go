package models

// Authority returns users authorities
type Authority struct {
	ID   uint   `json:"-" binding:"required"`
	Name string `json:"name" binding:"required"`
}
