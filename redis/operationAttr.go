package redis

import (
	"fmt"
	"time"
)

type OperationAttr struct {
	Name string
	Value interface{}
}

func NewOperationAttr(name string, value interface{}) *OperationAttr {
	return &OperationAttr{Name: name, Value: value}
}

type OperationAttrs []*OperationAttr
func(this OperationAttrs) Find(name string) *InterfaceResult {
	for _,attr := range this {
		if attr.Name == name {
			return NewInterfaceResult(attr.Value, nil)
		}
	}
	return NewInterfaceResult(nil, fmt.Errorf("OperationAttrs found error:%s",name))
}

type empty struct {}
const (
	ATTR_EXPIRE = "expire"	// 过期时间
	ATTR_NX = "nx"			// setnx
)
func WithExpire(t time.Duration) *OperationAttr {
	return NewOperationAttr(ATTR_EXPIRE, t)
}

func WithNx() *OperationAttr {
	return NewOperationAttr(ATTR_NX, empty{})
}
