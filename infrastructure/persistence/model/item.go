package model

type Item struct {
	Id        int `json:"id" gorm:"AUTO_INCREMENT" gorm:"primaryKey"`
	ProductId int `json:"product_id" gorm:"uniqueIndex;not null"`
}
