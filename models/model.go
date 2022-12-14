package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string
	Orders   []Order
}

type Order struct {
	gorm.Model
	UserID uint
	Price  float64
}
