package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Suggestion struct {
	Id          uint `gorm:"AUTO_INCREMENT" gorm:"primaryKey"`
	UUID        uuid.UUID
	UserId      uint
	Description string
	gorm.Model
}
