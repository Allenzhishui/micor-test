package conf

import (
	"fmt"
	"github.com/spf13/viper"
	"user/model"
)

func Init() {
	viper.AutomaticEnv()
	viper.SetConfigName("conf")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./conf/")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w \n", err))
	}
	model.Database()
}
