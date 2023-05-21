package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Nama  string `gorm:"not null"`
	Age   int    `gorm:"not null"`
	Phone string
}
