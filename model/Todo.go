package model

import "github.com/jinzhu/gorm"

type Todo struct {
	gorm.Model
	User string
	Todo string
}
