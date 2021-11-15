package redis

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"time"
)

const (
	SERILIZER_JSON = "json"
	SERILIZER_GOB = "gob"
)

type DBGetterFunc func() interface{}
type StringCache struct {
	Operation *StringOperation
	Expire time.Duration
	DBGetter DBGetterFunc
	Serilizer string
}

func NewStringCache(operation *StringOperation, expire time.Duration, serilizer string) *StringCache {
	return &StringCache{Operation: operation, Expire: expire, Serilizer: serilizer}
}


// 设置缓存
func(this *StringCache) SetCache(key string, value interface{}) {
	this.Operation.Set(key, value, WithExpire(this.Expire)).Unwrap()
}

// 读取缓存
func(this *StringCache) GetCache(key string) (ret interface{}) {
	obj := this.DBGetter()
	if this.Serilizer == SERILIZER_JSON {
		f := func() string {
			bs,err := json.Marshal(obj)
			if err != nil {
				return ""
			}
			return string(bs)
		}
		ret = this.Operation.Get(key).UnwrapOrElse(f)
		this.SetCache(key, ret)

	} else if this.Serilizer == SERILIZER_GOB {
		f := func() string {
			var buf = &bytes.Buffer{}
			enc := gob.NewEncoder(buf)
			if err := enc.Encode(obj); err != nil {
				return ""
			}
			return buf.String()
		}
		ret = this.Operation.Get(key).UnwrapOrElse(f)
		this.SetCache(key, ret)
	}

	return
}

func(this *StringCache) GetCacheForObject(key string,obj interface{})  interface{} {
	ret := this.GetCache(key)
	if ret == nil{
		return nil
	}

	if this.Serilizer == SERILIZER_JSON {
		err := json.Unmarshal([]byte(ret.(string)),obj)
		if err != nil {
			return nil
		}

	} else if this.Serilizer==SERILIZER_GOB {
		var buf = &bytes.Buffer{}
		buf.WriteString(ret.(string))
		dec:=gob.NewDecoder(buf)
		if dec.Decode(obj) != nil {
			return nil
		}
	}

	return nil
}