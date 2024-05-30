package data

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name           string
	Email          string `gorm:"unique"`
	Address        string
	PhoneNumber    string `gorm:"unique"`
	Pin            string
	Role           string
	ProfilePicture string
}
