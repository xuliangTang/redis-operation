package redis

import "time"

type OperationAttr struct {
	Name string
	Value interface{}
}

func NewOperationAttr(name string, value interface{}) *OperationAttr {
	return &OperationAttr{Name: name, Value: value}
}

type OperationAttrs []*OperationAttr
func(this OperationAttrs) Find(name string) interface{} {
	for _,attr := range this {
		if attr.Name == name {
			return attr.Value
		}
	}
	return nil
}

const (
	ATTR_EXPIRE = "expire"	// 过期时间
)
func WithExpire(t time.Duration) *OperationAttr {
	return NewOperationAttr(ATTR_EXPIRE, t)
}
