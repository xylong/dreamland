package db

import (
	"dreamland/pkg/model"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

var DB *gorm.DB

func InitDB() {
	db, err := gorm.Open(viper.GetString("database.driver"), fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		viper.GetString("database.user"),
		viper.GetString("database.password"),
		viper.GetString("database.host"),
		viper.GetString("database.name")))

	if err != nil {
		panic("connect mysql failed, " + err.Error())
	}

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(10)
	DB = db
	migration()
}

// 数据迁移
func migration() {
	if !DB.HasTable(&model.User{}) {
		DB.AutoMigrate(&model.User{})
	}
}

// CloseDB 关闭数据库
func CloseDB() {
	defer DB.Close()
}
