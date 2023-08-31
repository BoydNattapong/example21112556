package models

import (
	

	"gorm.io/gorm"
)

type Personservice struct {
	gorm.Model
	Cid string `gorm:"unique"`
	Fullname string 
	Choice1 string
	Choice2 string
	Choice3 string
	Choice4 string
	Choice5 string
	Score string
	Result_score string
}