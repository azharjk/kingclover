package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"kingclover.com/api/pkg/config"
	"kingclover.com/api/pkg/model"
)

var DB *gorm.DB

func InitializeDB() {
	db, err := gorm.Open(sqlite.Open(config.DatabaseName), &gorm.Config{})
	if err != nil {
		log.Panicln("Failed to connect to database")
	}

	err = db.AutoMigrate(&model.User{}, &model.AccessToken{}, &model.RefreshToken{})
	if err != nil {
		log.Panicln("Failed to migrate model")
	}

	DB = db
}
