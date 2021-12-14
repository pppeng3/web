package limit

import (
	"fmt"
	goredis "github.com/go-redis/redis"
	"time"
	"web/db/redis"

	"github.com/sirupsen/logrus"
)

type LimitMap struct {
	ttl int `json:"ttl"`
	num int `json:"num"`
}

var (
	limitMaps map[string]LimitMap
)

func init() {
	limitMaps = make(map[string]LimitMap, 100)
	limitMaps["ropz"] = LimitMap{
		ttl: 20,
		num: 1,
	}
	limitMaps["ropz2"] = LimitMap{
		ttl: 20,
		num: 1,
	}
}

// 统计时间内请求次数
func AddLimitData(typ, key string) int {
	redisClient := redis.Instance(15)
	defer redisClient.Close()
	_, ok := limitMaps[typ]
	if !ok {
		logrus.Errorf(`%s type is not exists`, typ)
		return -1
	}
	rKey := fmt.Sprintf("%s_%s", typ, key)
	err := redisClient.Get(rKey).Err()
	if err != nil {
		redisClient.Set(rKey, 1, time.Duration(limitMaps[typ].ttl)*time.Second)
		return 1
	} else {
		return int(redisClient.Incr(rKey).Val())
	}
}

// 查询请求次数是否超频,超频返回false
func HasOverflow(typ, key string) bool {
	redisClient := redis.Instance(15)
	defer redisClient.Close()
	_, ok := limitMaps[typ]
	if !ok {
		logrus.Errorf(`%s type is not exists`, typ)
	}
	rKey := fmt.Sprintf("%s_%s", typ, key)
	l, err := redisClient.Get(rKey).Int()
	if err == goredis.Nil {
		return true
	} else if err != nil {
		logrus.Errorf(`Get %s failed`, rKey)
		panic(err)
	}

	if l > limitMaps[typ].num {
		return false
	}
	return true
}

// 防刷是否超过频率(高并发)
func CheckUpperLimit(typ, key string) bool {
	redisClient := redis.Instance(15)
	defer redisClient.Close()
	_, ok := limitMaps[typ]
	if !ok {
		logrus.Errorf(`%s type is not exists`, typ)
		return false
	}
	rKey := fmt.Sprintf("%s_%s", typ, key)
	result := int(redisClient.Incr(rKey).Val())
	if result == 1 {
		redisClient.Expire(rKey, time.Duration(limitMaps[typ].ttl)*time.Second)
	} else if result > limitMaps[typ].num {
		return false
	}
	return true
}
