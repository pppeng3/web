package cache_service

import (
	"time"

	"github.com/patrickmn/go-cache"
)

var (
	def *cache.Cache
)

func init() {
	def = cache.New(time.Second*10, time.Minute*5)
}

func Instance() *cache.Cache {
	return def
}

// example cache_service.Instance().Get(key) cache_service.Instance().Set(k, v, d)
