package redis

import (
	"time"
)

type DBGetterFunc func() string
type StringCache struct {
	Operation *StringOperation
	Expire time.Duration
	DBGetter DBGetterFunc
}

func NewStringCache(operation *StringOperation, expire time.Duration) *StringCache {
	return &StringCache{Operation: operation, Expire: expire}
}

// 设置缓存
func(this *StringCache) SetCache(key string, value interface{}) {
	this.Operation.Set(key, value, WithExpire(this.Expire)).Unwrap()
}

// 读取缓存
func(this *StringCache) GetCache(key string) (ret interface{}) {
	ret = this.Operation.Get(key).UnwrapOrElse(this.DBGetter)
	this.SetCache(key, ret)
	return
}