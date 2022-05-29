package gedis

import "time"

const (
	ATTR_EXPR = "expr" //过期时间
)

type OperationAttr struct {
	Name string
	Value interface{}
}

type OperationAttrs []*OperationAttr

func (this OperationAttrs) Find(name string) interface{} {
	for _, attr := range this {
		if attr.Name == name {
			return attr.Value
		}
	}

	return nil
}

func WithExpire(t time.Duration) *OperationAttr {
	return &OperationAttr{Name: ATTR_EXPR, Value: t}
}