package main

import (
	//	"fmt"

	"example21112556/controllers"
	"example21112556/initializers"
	"example21112556/middleware"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	
)

func init() {

	initializers.LoadEnvVariable()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main() {

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true

	r := gin.Default()
	r.Use(cors.New(corsConfig))

	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.POST("/screening_covid", controllers.ScreeningCovid19)

	r.GET("/validate", middleware.RequireAuth, controllers.Validate)

	r.GET("/generatejwt", controllers.GenerateJWT)

	r.GET("/cardid/:cid", controllers.SelectIdCard)
	r.GET("/useremail/:email", controllers.GetUser)

	r.Run()
}
