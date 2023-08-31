package controllers

import (
	"example21112556/initializers"
	"example21112556/models"
	"fmt"
	

	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func GetUser(c *gin.Context) {

	email := c.Param("email")
	
	var users models.User
	
	initializers.DB.Where("email = ?", email).First(&users)
	
	//initializers.DB.Find(&patients,cid)
	//println("Hello")
	//c.JSON(200,gin.H{"data":patients})
	c.JSON(200,&users)
	}

func Signup(c *gin.Context) {
	//Get the email/pass off req body

	var body struct {
		Email    string
		Password string
		Cid      string
		Fullname string
		Position string
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	//Hash the password

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
		return

	}

	//Create the user
	//user := models.User{Email: body.Email, Password: body.Password, Cid: body.Cid, Fullname: body.Fullname, Position: body.Position}

	user := models.User{Email: body.Email, Password: string(hash), Cid: body.Cid, Fullname: body.Fullname, Position: body.Position}

	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})
		return
	}

	//Respond

	c.JSON(http.StatusOK, gin.H{})

}

func Login(c *gin.Context) {

	var body struct {
		Email    string
		Password string
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	//Look up requested user

	var user models.User
	initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	//Compare sent in pass with saved user pass hash

	//hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Password not match",
		})
		return
	}

	//Generate a jwt token

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   user.ID,
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECERET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token",
		})
		return
	}

	//SHOW TOKEN
	/*	c.JSON(http.StatusOK,gin.H{
		"token":tokenString,
	})  */

	//send it back
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{})

}

func Validate(c *gin.Context) {
	//	user, _ := c.Get("user")

	tokenString, _ := c.Cookie("Authorization")

	//fmt.Println(user)
	c.JSON(http.StatusOK, gin.H{
		"message": "I'm logged in ",
		"token":   tokenString,
	})
}

func GenerateJWT(c *gin.Context) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo":  "bar",
		"name": "KONG",

		//	"nbf": time.Date(2022, 12, 10, 12, 0, 0, 0, time.UTC).Unix(),
		"nbf": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECERET")))

	fmt.Println(tokenString, err)

}
