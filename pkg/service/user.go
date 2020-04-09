package service

import (
	"dreamland/pkg/dao"
	"dreamland/pkg/model"
	"dreamland/pkg/validate"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Login(req *validate.LoginRequest) (user *model.User, err error)
	Register(req *validate.RegisterRequest) (string, error)
}

func NewUserService() UserService {
	return &User{dao.NewUser()}
}

type User struct {
	user dao.User
}

func (u *User) Login(req *validate.LoginRequest) (user *model.User, err error) {
	user, err = u.user.Find(&model.User{
		Email: req.Email,
	})
	if err != nil {
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, errors.New("密码错误")
	}
	return
}

func (u *User) Register(req *validate.RegisterRequest) (string, error) {
	if u.user.IsEmailExit(req.Email) {
		return "", errors.New("邮箱已被注册")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	_, err = u.user.Insert(&model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hash),
	})
	return "token", err
}
