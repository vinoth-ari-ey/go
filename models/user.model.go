package models

import "gorm.io/gorm"

// User represents a user in the database.
type User struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email" gorm:"unique"`
	Age   int    `json:"age"`
}
