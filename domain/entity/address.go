package entity

import "github.com/google/uuid"

type Address struct {
	UUID         uuid.UUID `json:"uuid"`
	UserUUID     uuid.UUID `json:"userUUID"`
	Street       string    `json:"street"`
	Neighborhood string    `json:"neighborhood"`
	Number       string    `json:"number"`
	City         string    `json:"city"`
	Complement   string    `json:"complement"`
	Landmark     string    `json:"landmark"`
}
