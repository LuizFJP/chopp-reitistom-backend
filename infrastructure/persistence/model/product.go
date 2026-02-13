package model

import "github.com/google/uuid"

type Product struct {
	Id          uint      `gorm:"AUTO_INCREMENT" gorm:"primaryKey"`
	UUID        uuid.UUID `gorm:"not null"`
	Name        string    `gorm:"not null"`
	Price       float64   `gorm:"not null"`
	CategoryId  uint
	Ratings     []*Rating     `gorm:"foreignKey:ProductId"`
	Ingredients []*Ingredient `gorm:"foreignKey:ProductId"`
	Items       []*Item       `gorm:"foreignKey:ProductId"`
	Users       []*User       `gorm:"many2many:favorite_product"`
	Orders      []*Order      `gorm:"many2many:product_order"`
}
