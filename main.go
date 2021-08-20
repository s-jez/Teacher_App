package main

import (
	// Nie ma sensu dawać odnośnika do twojego projektu do github.com
	// Jeżeli tutaj chcesz zaimportować jakiś plik ze swojego projektu, musisz to robić lokalnie, czyli tak jak ci zmieniłem :)
	// "github.com/Stachowsky/Teacher_App/config"
	// "github.com/Stachowsky/Teacher_App/routers"
	"Stachowsky/Teacher_App/config"
	"Stachowsky/Teacher_App/routers"
	"fmt"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Zmienna do errorów żeby połączenie do bazy nie szalało
var err error

func main() {
	// Nie potrzebne, już ładujesz templatki w Routersach
	// config.LoadTemplate()

	// Create students.db file
	config.CreateFile()

	// W taki sposób nie będziesz w stanie przypisać połączenia do bazy do zmiennej globalnej aby używać jej w każdym możliwym pliku w projekcie
	// config.Connection()

	// Połączenie do bazy jest główną funkcjonalnością serwera, musi być ustawiona w main żeby nie szukać po wszystkich plikach gdzie co jest
	config.DB, err = gorm.Open(sqlite.Open("students.db"), &gorm.Config{Logger: logger.Default.LogMode(logger.Error)})
	if err != nil {
		fmt.Println("statuse: ", err)
		os.Exit(1)
	}

	// Dzięki takiemu zabiegowi wiemy co i gdzie się odpala. Nie musimy grzebać po wszystkich plikach żeby zobaczyć gdzie jest ustawione odpalanie serwerka
	// Main jest idealnym miejscem do odpalenia głównych funkcjonalności, między innymi routera
	Router := routers.CreateUrlMappings()
	// Można ładnie określić port
	Router.Run(":8080")
}
