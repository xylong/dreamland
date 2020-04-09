package dao

import (
	"dreamland/pkg/db"
	"dreamland/pkg/model"
)

type User interface {
	Insert(user *model.User) (id uint, err error)
	IsEmailExit(email string) bool
}

func NewUser() User {
	return &user{}
}

type user struct {
}

func (u *user) Insert(user *model.User) (id uint, err error) {
	err = db.DB.Create(user).Error
	if err != nil {
		return 0, err
	}
	return user.ID, nil
}

func (u *user) IsEmailExit(email string) bool {
	var user model.User
	db.DB.Select("id").Where("email=?", email).First(&user)
	return user.ID > 0
}
