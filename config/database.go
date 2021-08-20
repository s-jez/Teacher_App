package config

import (
	"gorm.io/gorm"
)

// Zmienna globalna, pozwoli ci na połączenie się do bazy w każdym folderze po zaimportowaniu /Stachowsky/config :)
var DB *gorm.DB
