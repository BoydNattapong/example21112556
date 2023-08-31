package initializers

import (

	"log"
	"os"


	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	
)



var DB *gorm.DB

func ConnectToDB() {


    var err error
//	dsn := "sa:sa@tcp(127.0.0.1:3306)/dtcdb?charset=utf8mb4&parseTime=True&loc=Local"
    dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	  }
	

}