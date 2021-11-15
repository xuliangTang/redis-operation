package cachePool

import (
	"goRedis/redis"
	"sync"
	"time"
)

var NewsCachePool *sync.Pool

func init()  {
	NewsCachePool = &sync.Pool{New: func() interface{} {
		return redis.NewStringCache(redis.NewStringOperation(), time.Second*30,
			redis.SERILIZER_GOB, redis.NewCrossPolicy("^news\\d{1,6}$", time.Second*10))
	}}
}

func NewsCache() *redis.StringCache {
	return NewsCachePool.Get().(*redis.StringCache)
}

func NewsCacheRelease(cache *redis.StringCache)  {
	NewsCachePool.Put(cache)
}