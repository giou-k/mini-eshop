package model

import "github.com/jinzhu/gorm"

type Category struct {
	gorm.Model

	Title    string    `json:"title"`
	Position int       `json:"position"` //(for ordering)
	ImageURL string    `json:"image_url"`
	Product  []Product `json:"product,omitempty"` //gorm:"foreignkey:CategoryID
}
