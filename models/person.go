package models

import "github.com/jinzhu/gorm"

type Person struct {
	gorm.Model
	Name     string `json:"name"`
	LastName string `json:"lastname"`
	Age      uint8  `json:"age"`
}
