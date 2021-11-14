package redis

import (
	"context"
	"time"
)

type StringOperation struct {
	ctx context.Context
}

func NewStringOperation() *StringOperation {
	return &StringOperation{ctx: context.Background()}
}

func(this *StringOperation) Set(key string, val string, attrs ...*OperationAttr) *StringResult {
	exp := OperationAttrs(attrs).Find(ATTR_EXPIRE)
	if exp == nil {
		exp = time.Second*0
	}
	return NewStringResult(Redis().Set(this.ctx, key, val, exp.(time.Duration)).Result())
}

func(this *StringOperation) Get(key string) *StringResult {
	return NewStringResult(Redis().Get(this.ctx, key).Result())
}

func(this *StringOperation) MGet(keys ...string) *SliceResult {
	return NewSliceResult(Redis().MGet(this.ctx, keys...).Result())
}