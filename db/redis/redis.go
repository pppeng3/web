package redis

import (
	"web/config"

	"github.com/go-redis/redis"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

var (
	redisClient *redis.Client
	err         error
)

func Instance(db int) *redis.Client {
	config.Init()
	conf := config.GetConfig()
	redisClient = redis.NewClient(&redis.Options{
		Addr:     conf.Redis.Addr,
		Password: conf.Redis.Password,
		DB:       db,
	})
	_, err = redisClient.Ping().Result()
	if err != nil {
		logrus.Error(errors.WithStack(err))
		panic(err)
	}
	return redisClient
}
