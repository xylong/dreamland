package service

import (
	"dreamland/pkg/dao"
	"dreamland/pkg/dto"
	"dreamland/pkg/model"
	"dreamland/pkg/util"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"strconv"
)

type UserService interface {
	Login(req *dto.LoginRequest) (token string, err error)
	Register(req *dto.RegisterRequest) error
	RefreshToken(oldToken string) (string, error)
}

func NewUserService() UserService {
	return &User{dao.NewUser()}
}

type User struct {
	user dao.User
}

func (u *User) Login(req *dto.LoginRequest) (token string, err error) {
	user, err := u.user.Find(&model.User{
		Email: req.Email,
	})
	if err != nil {
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return "", errors.New("密码错误")
	}
	id := strconv.Itoa(int(user.ID))
	token, err = util.NewJWT().Generate(&util.Claims{
		ID:    id,
		Name:  user.Name,
		Email: user.Email,
	})
	return
}

func (u *User) Register(req *dto.RegisterRequest) error {
	if u.user.IsEmailExit(req.Email) {
		return errors.New("邮箱已被注册")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	_, err = u.user.Insert(&model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hash),
	})
	return err
}

func (u *User) RefreshToken(oldToken string) (string, error) {
	return util.NewJWT().Refresh(oldToken)
}
