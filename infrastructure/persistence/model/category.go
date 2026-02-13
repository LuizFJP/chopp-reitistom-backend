package model

import "github.com/google/uuid"

type Category struct {
	Id      uint      `gorm:"AUTO_INCREMENT" gorm:"primaryKey"`
	UUID    uuid.UUID `gorm:"not null"`
	Name    string    `gorm:"uniqueIndex;not null"`
	Product []Product `gorm:"foreignKey:CategoryId"`
}
