package models

import (
	"time"
)

type User struct {
	ID        uint      `json:"id",gorm:"primary_key",sql:"type:int unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Name      string    `json:"name"`
	LastName  string    `json:"lastName"`
	Age       int       `json:"age"`
}
