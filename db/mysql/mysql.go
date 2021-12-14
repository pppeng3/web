package mysql

import (
	"fmt"
	"web/config"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var (
	mysqlDB *gorm.DB
	err     error
)

const (
	base string = "%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true"
)

func Instance() *gorm.DB {
	config.Init()
	conf := config.GetConfig()
	dsn := fmt.Sprintf(base, conf.Mysql.Username, conf.Mysql.Password, conf.Mysql.Host, conf.Mysql.Port, conf.Mysql.Dbname)
	mysqlDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:                 logger.Default.LogMode(logger.Info),
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
			NoLowerCase:   false,
		},
	})
	if err != nil {
		logrus.Error(errors.WithStack(err))
		panic(err)
	}
	return mysqlDB.Debug()
}
