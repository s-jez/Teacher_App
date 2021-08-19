package config

import (
	"github.com/Stachowsky/Teacher_App/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connection() {
	// Connect to SQLite database...
	db, err := gorm.Open(sqlite.Open("students.db"), &gorm.Config{Logger: logger.Default.LogMode(logger.Error)})
	if err != nil {
		panic("Failed connection to SQLite database!")
	}
	// Create table in database by Student struct
	db.AutoMigrate(&models.Student{})
}
