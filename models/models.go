package model

import "gorm.io/gorm"

type Album struct {
	gorm.Model
	Title string  `json:"title"`
	Actor string  `json:"actor"`
	Price float64 `json:"price"`
}
