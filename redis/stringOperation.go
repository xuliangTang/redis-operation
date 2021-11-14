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

func(this *StringOperation) Set(key string, val interface{}, attrs ...*OperationAttr) *InterfaceResult {
	exp := OperationAttrs(attrs).Find(ATTR_EXPIRE).UnwrapOr(time.Second*0).(time.Duration)
	nx := OperationAttrs(attrs).Find(ATTR_NX).UnwrapOr(nil)

	if nx != nil {
		return NewInterfaceResult(Redis().SetNX(this.ctx, key, val, exp).Result())
	}

	// nx 和 xx 同时只能出现一个
	xx := OperationAttrs(attrs).Find(ATTR_XX).UnwrapOr(nil)
	if xx != nil {
		return NewInterfaceResult(Redis().SetXX(this.ctx, key, val, exp).Result())
	}

	return NewInterfaceResult(Redis().Set(this.ctx, key, val, exp).Result())
}

func(this *StringOperation) Get(key string) *StringResult {
	return NewStringResult(Redis().Get(this.ctx, key).Result())
}

func(this *StringOperation) MGet(keys ...string) *SliceResult {
	return NewSliceResult(Redis().MGet(this.ctx, keys...).Result())
}