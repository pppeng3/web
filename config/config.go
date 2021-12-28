package config

import (
	"web/log"

	"github.com/spf13/viper"
)

func Init() {
	viper.AddConfigPath("../conf")
}

func GetConfig() Conf {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Error(err.Error())
		panic("viper readInitConfig Error")
	}
	var conf Conf
	if err := viper.Unmarshal(&conf); err != nil {
		log.Error(err.Error())
	}
	return conf
}
