package models

import "gorm.io/gorm"

type ScreeningCovid19 struct {
	gorm.Model
	Staff_Cid string  `json:"staff_cid"`
	Staff_Fullname string  `json:"staff_fullname"`
	Answer_Cid string  `json:"answer_cid"`
	Answer_Fullname string  `json:"answer_fullname"`
	Score  string  `json:"score"`
	Score_Result string  `json:"score_result"`
	
}