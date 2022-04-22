package entity

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Judul  string
	Author string
	Owner  uint
}
