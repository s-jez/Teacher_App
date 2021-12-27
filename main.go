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

//TODO dev branch:
// 1 (frontend update)
// use header json instead of form value
// add to update current record values
// 2 (dev branch)
// JWT - Authentication Refresh Token and Access Token
// Create account (register/login)
// create tables in db and struct
// add endpoints to backend
// roles: admin (all permissions, view/edit/delete) / dev (view/edit)
// anonymous (unlogged) users can only view users
// authentication JWT
// user login have access_token and refresh_token from backend
// for next request from backend user send in header Autherization a access token
// backend verify access token is avaible and for access token verify user (jwt - user_id)
// for user_id we are searching if current user have role to do request from backend
// if access_token expires it have to be refreshed from refresh_token
// if refresh_token too expires, user must login again to account
// authentication user in Middleware
// at the end we can do unit tests
