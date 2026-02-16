package entity

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Status string

const (
	StatusCreated        Status = "CRIADO"
	StatusConfirmed      Status = "CONFIRMADO"
	StatusCanceled       Status = "CANCELADO"
	StatusSent           Status = "ENVIADO"
	StatusWaitingPayment Status = "AGUARDANDO_PAGAMENTO"
)

func (s Status) Valid() bool {
	switch s {
	case StatusCreated, StatusConfirmed, StatusCanceled, StatusSent, StatusWaitingPayment:
		return true
	default:
		return false
	}
}

func ParseStatus(v string) (Status, error) {
	s := Status(v)
	if !s.Valid() {
		return "", fmt.Errorf("invalid status: %q", v)
	}
	return s, nil
}

type Order struct {
	UUID          uuid.UUID   `json:"uuid"`
	UserUUID      uuid.UUID   `json:"userUUID"`
	ProductsUUID  []uuid.UUID `json:"productsUuid"`
	Status        Status      `json:"status"`
	CreatedAt     time.Time   `json:"createdAt"`
	ConfirmedAt   time.Time   `json:"confirmedAt"`
	SentAt        time.Time   `json:"sentAt"`
	CanceledAt    time.Time   `json:"canceledAt"`
	PaymentMethod string      `json:"paymentMethod"`
	StatusPayment string      `json:"statusPayment"`
}
