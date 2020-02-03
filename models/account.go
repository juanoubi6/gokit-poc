package models

import (
	"time"
)

type Account struct {
	ID        uint      `json:"id",gorm:"primary_key",sql:"type:int unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Email     string    `json:"email" gorm:"unique_index:idx_email"`
	Password  string    `json:"-"`
}
