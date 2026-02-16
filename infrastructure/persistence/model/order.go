package model

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	Id            uint      `gorm:"AUTO_INCREMENT" gorm:"primaryKey"`
	UUID          uuid.UUID `gorm:"not null"`
	UserId        uint
	Products      []*Product `gorm:"many2many:product_order"`
	Status        string
	CreatedAt     time.Time
	ConfirmedAt   time.Time
	SentAt        time.Time
	CanceledAt    time.Time
	PaymentMethod string
	StatusPayment string
}
