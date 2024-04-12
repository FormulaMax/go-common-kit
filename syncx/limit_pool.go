package syncx

import "sync/atomic"

// LimitPool 是对 Pool 的简单封装允许用户通过控制一段时间内对Pool的令牌申请次数来间接控制Pool中对象的内存总占用量
type LimitPool[T any] struct {
	pool   *Pool[T]
	tokens *atomic.Int32
}

func NewLimitPool[T any](maxTokens int, factory func() T) *LimitPool[T] {
	var tokens atomic.Int32
	tokens.Add(int32(maxTokens))
	return &LimitPool[T]{
		pool:   NewPool[T](factory),
		tokens: &tokens,
	}
}

// Get 取出一个元素
func (lp *LimitPool[T]) Get() T {
	if lp.tokens.Add(-1) < 0 {
		lp.tokens.Add(1)
		var zero T
		return zero
	}
	return lp.pool.Get()
}

// Put 放回去一个元素
func (lp *LimitPool[T]) Put(t T) {
	lp.pool.Put(t)
	lp.tokens.Add(1)
}
