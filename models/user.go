package models

import (
	"time"
)

type User struct {
	ID        uint `gorm:"primary_key",sql:"type:int unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string `json:"name"`
	LastName  string `json:"lastName"`
	Age       int    `json:"age"`
}
