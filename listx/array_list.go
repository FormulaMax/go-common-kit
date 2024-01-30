package listx

import (
	"github.com/FormulaMax/go-common-kit/internal/errs"
	"github.com/FormulaMax/go-common-kit/internal/slice"
)

var (
	_ List[any] = &ArrayList[any]{}
)

type ArrayList[T any] struct {
	vals []T
}

func NewArrayList[T any](cap int) *ArrayList[T] {
	return &ArrayList[T]{
		vals: make([]T, 0, cap),
	}
}

func NewArrayListOf[T any](ts []T) *ArrayList[T] {
	return &ArrayList[T]{
		vals: ts,
	}
}

func (a *ArrayList[T]) Get(index int) (T, error) {
	l := a.Len()
	if index < 0 || index >= l {
		var zero T
		return zero, errs.NewErrIndexOutOfRange(l, index)
	}
	return a.vals[index], nil
}

func (a *ArrayList[T]) Append(ts ...T) error {
	a.vals = append(a.vals, ts...)
	return nil
}

func (a *ArrayList[T]) Add(index int, t T) error {
	if index < 0 || index > len(a.vals) {
		return errs.NewErrIndexOutOfRange(len(a.vals), index)
	}
	a.vals = append(a.vals, t)
	copy(a.vals[index+1:], a.vals[index:])
	a.vals[index] = t
	return nil
}

func (a *ArrayList[T]) Set(index int, t T) error {
	length := len(a.vals)
	if index < 0 || index > length {
		return errs.NewErrIndexOutOfRange(length, index)
	}
	a.vals[index] = t
	return nil
}

// Delete 方法会在必要的时候引起缩容，其缩容规则是：
// - 如果容量 > 2048，并且长度小于容量一半，那么就会缩容为原本的 5/8
// - 如果容量 (64, 2048]，如果长度是容量的 1/4，那么就会缩容为原本的一半
// - 如果此时容量 <= 64，那么我们将不会执行缩容。在容量很小的情况下，浪费的内存很少，所以没必要消耗 CPU去执行缩容
func (a *ArrayList[T]) Delete(index int) (T, error) {
	res, t, err := slice.Delete(a.vals, index)
	if err != nil {
		return t, err
	}
	a.vals = res
	a.shrink()
	return t, nil
}

func (a *ArrayList[T]) shrink() {
	a.vals = slice.Shrink[T](a.vals)
}

func (a *ArrayList[T]) Len() int {
	return len(a.vals)
}

func (a *ArrayList[T]) Cap() int {
	return cap(a.vals)
}

func (a *ArrayList[T]) Range(fn func(index int, t T) error) error {
	for key, value := range a.vals {
		e := fn(key, value)
		if e != nil {
			return e
		}
	}
	return nil
}

func (a *ArrayList[T]) AsSlice() []T {
	res := make([]T, len(a.vals))
	copy(res, a.vals)
	return res
}
