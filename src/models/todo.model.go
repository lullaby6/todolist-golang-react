package models

import "gorm.io/gorm"

type ToDo struct {
	gorm.Model
	ID    uint   `gorm:"not null;primaryKey;autoIncrement;unique_index"`
	Title string `gorm:"not null"`
	Done  bool   `gorm:"not null;default:false"`
}
