package models

import "gorm.io/gorm"

type Patient struct {
	gorm.Model
	Cid      string `gorm:"unique" json:"cid"`
	Fullname string  `json:"fullname"`
}
