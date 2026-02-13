package model

import "github.com/google/uuid"

type Address struct {
	Id           uint      `gorm:"AUTO_INCREMENT" gorm:"primaryKey"`
	UUID         uuid.UUID `gorm:"not null"`
	Street       string    `gorm:"not null"`
	Neighborhood string    `gorm:"not null"`
	Number       string
	City         string `gorm:"not null"`
	Complement   string
	Landmark     string
	UserId       uint
	User         User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL" gorm:"foreignKey:UserId"`
}
