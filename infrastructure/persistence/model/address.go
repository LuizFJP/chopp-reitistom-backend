package model

import "github.com/google/uuid"

type Address struct {
	Id           int       `json:"id" gorm:"AUTO_INCREMENT" gorm:"primaryKey"`
	UserId       int       `json:"user_id" gorm:"uniqueIndex;not null"`
	UUID         uuid.UUID `json:"uuid"`
	Street       string    `json:"street"`
	Neighborhood string    `json:"neighborhood"`
	Number       string    `json:"number"`
	City         string    `json:"city"`
	Complement   string    `json:"complement"`
	Landmark     string    `json:"landmark"`
}
