package models

import "gorm.io/gorm"

type ToDo struct {
	gorm.Model
	Title string
	Done  bool
}
