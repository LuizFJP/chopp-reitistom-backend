package entity

import (
	"fmt"

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
	OrderUUID   uuid.UUID `json:"orderUUID"`
	UserUUID    uuid.UUID `json:"userUUID"`
	ProductUUID uuid.UUID `json:"productUUID"`
	Status      Status    `json:"status"`
}
