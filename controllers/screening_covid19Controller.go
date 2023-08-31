package controllers

import (
	"example21112556/initializers"
	"example21112556/models"
	

	"net/http"
	

	"github.com/gin-gonic/gin"


)

func ScreeningCovid19(c *gin.Context) {
	//Get the email/pass off req body


	var body struct {
		Staff_Cid string
		Staff_Fullname    string
		Answer_Cid      string
		Answer_Fullname string
		Score string
		Score_Result string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}



	screening_covid19 := models.ScreeningCovid19{Staff_Cid: body.Staff_Cid, Staff_Fullname: body.Staff_Fullname , Answer_Cid: body.Answer_Cid, Answer_Fullname: body.Answer_Fullname, Score: body.Score, Score_Result: body.Score_Result}

	result := initializers.DB.Create(&screening_covid19)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create screening_covid19",
		})
		return
	}

	//Respond

	c.JSON(http.StatusOK, gin.H{})

}