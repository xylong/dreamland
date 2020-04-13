package dao

import (
	"dreamland/config"
	"dreamland/pkg/db"
	"dreamland/pkg/model"
	"dreamland/pkg/util"
	"testing"
)

func init() {
	config.InitConfig()
	db.InitDB()
}

func TestUser_Insert(t *testing.T) {
	user := NewUser()
	id, err := user.Insert(&model.User{
		Name:     util.RandomString(10),
		Avatar:   "",
		Email:    "123456@qq.com",
		Password: "123456",
	})
	if err != nil {
		panic(err)
	}
	t.Logf("id:%d\n", id)
}
