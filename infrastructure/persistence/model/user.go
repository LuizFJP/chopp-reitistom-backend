package model

import (
	"chopp-reitistom-backend/domain/entity"

	"github.com/google/uuid"
)

type User struct {
	Id       int            `json:"id" gorm:"AUTO_INCREMENT" gorm:"primaryKey"`
	Name     string         `json:"name"`
	UUID     uuid.UUID      `json:"UUID"`
	Email    string         `json:"email"`
	Password string         `json:"password"`
	Address  entity.Address `json:"address" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
