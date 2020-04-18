package config

import (
	"github.com/spf13/viper"
	"os"
)

func InitConfig() {
	workspace, _ := os.Getwd()
	viper.SetConfigName("app")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workspace + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
