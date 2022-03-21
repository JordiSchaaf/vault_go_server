package utils

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"vault/server/models"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("../vault_db.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to create a connection to database")
	}

	database.AutoMigrate(&models.User{})
	println("Database connected!")
	DB = database
}

func DisconnectDatabase() {
	database, err := DB.DB()
	if err != nil {
		panic("Failed to close connection from database")
	}
	database.Close()
}
