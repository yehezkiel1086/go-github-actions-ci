package model

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Name string `json:"name" gorm:"size:255;not null"`
	Email string `json:"email" gorm:"size:255;not null;unique"`
	Password string `json:"password" gorm:"size:255;not null"`
}