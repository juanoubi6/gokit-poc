package models

import (
	"time"
)

type Account struct {
	ID        uint `gorm:"primary_key",sql:"type:int unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Email     string `gorm:"unique_index:idx_email",json:"email"`
	Password  string `json:"-"`
}
