package config

import (
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func Init() {
	viper.AddConfigPath("../../conf")
}

func GetConfig() Conf {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		logrus.Error(errors.WithStack(err))
		panic("viper readInitConfig Error")
	}
	var conf Conf
	if err := viper.Unmarshal(&conf); err != nil {
		logrus.Error(errors.WithStack(err))
	}
	return conf
}
