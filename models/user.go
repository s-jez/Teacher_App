package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"Name"`
	Pass string `json:"Pass"`
	Email    string `gorm:"type:varchar(64);unique_index"`
}
