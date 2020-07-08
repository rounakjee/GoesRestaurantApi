package models

import (
	"github.com/jinzhu/gorm"
)


type User struct{
	gorm.Model
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Email string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	HashPassword string `json:"HashPassword"`
	Password string `gorm:"-" json:"password"`
	Role string `json:"role"`
}

