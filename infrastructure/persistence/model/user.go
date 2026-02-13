package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id               uint         `gorm:"AUTO_INCREMENT" gorm:"primaryKey"`
	Name             string       `gorm:"not null"`
	UUID             uuid.UUID    `gorm:"not null"`
	Email            string       `gorm:"not null"`
	Password         string       `gorm:"not null"`
	Role             string       `gorm:"not null"`
	Suggestion       []Suggestion `gorm:"foreignKey:UserId"`
	FavoriteProducts []*Product   `gorm:"many2many:favorite_product"`
	Orders           []*Order     `gorm:"foreignKey:UserId"`
}
