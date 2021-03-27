package utils

import (
	"boj-garden/models"
	"gorm.io/gorm"
)

func MigrateDatabase(db *gorm.DB) {
	err := db.AutoMigrate(&models.User{}, &models.Submission{})

	if err != nil {
		panic(err.Error())
	}
}
