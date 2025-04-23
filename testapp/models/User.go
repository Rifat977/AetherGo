package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name  string
	Email string `gorm:"unique"`
}

func GetAll() []interface{} {
	return []interface{}{
		&User{},
	}
}
