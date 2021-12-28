package main

import (
	"Stachowsky/Teacher_App/config"
	"Stachowsky/Teacher_App/models"
	"Stachowsky/Teacher_App/routers"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var err error

func main() {
	// Create database file
	config.CreateDBFile()
	// Create env file
	config.CreateEnvFile()
	// Load .env file
	config.LoadEnv()
	// Database connection
	config.DB, err = gorm.Open(sqlite.Open("students.db"), &gorm.Config{Logger: logger.Default.LogMode(logger.Error)})
	if err != nil {
		fmt.Println(err)
		return
	}
	// Create tables in database
	config.DB.AutoMigrate(&models.Student{}, &models.User{})
	// Initalize routers
	r := routers.CreateUrlMappings()
	// Run server
	r.Run(":8080")
}
