package model

type Product struct {
	Id         int     `json:"id" gorm:"AUTO_INCREMENT" gorm:"primaryKey"`
	CategoryId int     `json:"category_id" gorm:"uniqueIndex;not null"`
	Name       string  `json:"name"`
	Item       Item    `json:"item"`
	Price      float64 `json:"price"`
}
