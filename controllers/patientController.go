package controllers

import (
	"example21112556/initializers"
	"example21112556/models"

	"github.com/gin-gonic/gin"
)


func SelectIdCard(c *gin.Context) {

cid := c.Param("cid")

var patients models.Patient

initializers.DB.Where("cid = ?", cid).First(&patients)

//initializers.DB.Find(&patients,cid)
//println("Hello")
//c.JSON(200,gin.H{"data":patients})
c.JSON(200,&patients)



}