package main

import (
	"github.com/Stachowsky/Teacher_App/config"
	"github.com/Stachowsky/Teacher_App/routers"
)

func main() {
	// Load templates from path.
	config.LoadTemplate()
	// Create students.db file
	config.CreateFile()
	// Connection to the database
	config.Connection()
	// Setup routers and run server
	routers.SetupRouters()
}
