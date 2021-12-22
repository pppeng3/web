package redis

import (
	"web/config"

	"web/log"

	"github.com/go-redis/redis"
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
		log.Fatal(err.Error())
		panic(err)
	}
	return redisClient
}
