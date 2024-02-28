package initializers

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error

	dsn := os.Getenv("DBCONN")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})


	if err != nil {
		panic("Error connecting to database")
	}
}