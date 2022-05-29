package gedis

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

func (this *StringOperation) Get(key string) *StringResult {
	return NewStringResult(Redis().Get(this.ctx, key).Result())
}

func (this *StringOperation) MGet(keys ...string) *SliceResult {
	return NewSliceResult(Redis().MGet(this.ctx, keys...).Result())
}

func (this *StringOperation) Set(key string, value interface{}, attrs ...*OperationAttr) *StringResult {
	exp := OperationAttrs(attrs).Find(ATTR_EXPR)
	if exp == nil {
		exp = time.Second * 0
	}
	return NewStringResult(Redis().Set(this.ctx, key, value, exp.(time.Duration)).Result())
}

