package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name  string
	Email string `gorm:"unique"`
}

type Post struct {
	gorm.Model
	Title   string
	Content string
	UserID  uint
}

func GetAll() []interface{} {
	return []interface{}{
		&User{},
		&Post{},
	}
}
