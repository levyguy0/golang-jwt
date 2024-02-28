package initializers

import (
	"projects/golang-jwt/models"
)

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}