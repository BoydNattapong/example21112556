package initializers

import "example21112556/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Patient{})
	DB.AutoMigrate(&models.Personservice{})
	DB.AutoMigrate(&models.Login{})
	DB.AutoMigrate(&models.ScreeningCovid19{})
}