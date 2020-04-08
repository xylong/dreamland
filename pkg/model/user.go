package model

import "github.com/jinzhu/gorm"

// User 用户
type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(20);not null;"`
	Avatar   string `gorm:"type:varchar(100)"`
	Email    string `gorm:"type:varchar(100);not null;unique"`
	Password string `gorm:"type:varchar(255)"`
}
