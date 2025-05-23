package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"unique;not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
	IsAdmin  bool   `gorm:"default:false" json:"is_admin"`
	Orders   []Order
}
