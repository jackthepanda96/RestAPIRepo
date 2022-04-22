package entity

import "gorm.io/gorm"

type Pegawai struct {
	gorm.Model
	Nama     string
	HP       int
	Password string
	Books    []Book `gorm:"foreignKey:Owner"`
}
