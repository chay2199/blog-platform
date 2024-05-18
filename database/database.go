package database

import (
	"log"

	"blog-platform/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	var err error
	DB, err = gorm.Open(sqlite.Open("/app/database/blog.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}
	DB.AutoMigrate(&models.Post{})
	DB.AutoMigrate(&models.User{})
}
