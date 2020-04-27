package model

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model

	CategoryID uint `gorm:"index"` //`sql:"index"` //gorm:"column:category_id"` // json:"category_id

	Title       string  `json:"title"`
	ImageURL    string  `json:"image_url"`
	Price       float32 `json:"price"`
	Description string  `json:"description"`
}
