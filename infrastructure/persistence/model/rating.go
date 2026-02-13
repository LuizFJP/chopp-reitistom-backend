package model

import "github.com/google/uuid"

type Rating struct {
	Id        uint      `gorm:"AUTO_INCREMENT" gorm:"primaryKey"`
	UUID      uuid.UUID `gorm:"not null"`
	ProductId uint
	Comment   string `gorm:"not null"`
	Grade     uint   `gorm:"not null"`
}
