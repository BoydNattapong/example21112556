package models

import "gorm.io/gorm"

type Login struct {
	gorm.Model
	Email string `gorm:"unique"  json:"email"`
	Password string  `json:"password"`
	Cid string  `json:"cid"`
	Fullname string  `json:"fullname"`
	Position  string  `json:"position"`
}