package model

import "github.com/google/uuid"

type Ingredient struct {
	Id        uint      `gorm:"AUTO_INCREMENT" gorm:"primaryKey"`
	UUID      uuid.UUID `gorm:"not null"`
	ProductId uint
	Name      string `gorm:"not null"`
}
