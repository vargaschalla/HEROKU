package models

import (
	"gorm.io/gorm"
)

type Person struct {
	gorm.Model
	Name string `json:"name"`
	Age  string `json:"age"`
}
