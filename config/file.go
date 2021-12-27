package config

import (
	"log"
	"os"
)

func CreateDBFile() {
	// Create file students.db
	f, err := os.OpenFile("students.db", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	f.Close()
}
func CreateEnvFile() {
	// Create .env file
	f, err := os.OpenFile(".env", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	_, err = f.Write([]byte("ACCESS_SECRET_TOKEN= \nREFRESH_SECRET_TOKEN="))
	if err != nil {
		log.Fatal(err)
	}
	f.Close()
}
