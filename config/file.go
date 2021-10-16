package config

import (
	"log"
	"os"
)

func CreateFile() {
	// Create file students.db
	f, err := os.OpenFile("students.db", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	f.Close()
}
