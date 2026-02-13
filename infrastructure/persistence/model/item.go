package model

type Item struct {
	Id        uint `gorm:"AUTO_INCREMENT" gorm:"primaryKey"`
	ProductId uint
}
