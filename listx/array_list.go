package listx

import (
	"github.com/FormulaMax/go-common-kit/internal/errs"
	"github.com/FormulaMax/go-common-kit/internal/slice"
)

type ArrayList[T any] struct {
	vals []T
}

func NewArrayList[T any](cap int) *ArrayList[T] {
	return &ArrayList[T]{
		vals: make([]T, 0, cap),
	}
}

func NewArrayListOf[T any](src []T) *ArrayList[T] {
	return &ArrayList[T]{
		vals: src,
	}
}

func (a *ArrayList[T]) Get(index int) (t T, e error) {
	l := a.Len()
	if index < 0 || index >= l {
		return t, errs.NewErrIndexOutOfRange(l, index)
	}
	return a.vals[index], nil
}

func (a *ArrayList[T]) Add(index int, t T) (err error) {
	a.vals, err = slice.Add[T](a.vals, t, index)
	return
}

func (a *ArrayList[T]) Append(ts ...T) error {
	a.vals = append(a.vals, ts...)
	return nil
}

func (a *ArrayList[T]) Set(index int, t T) error {
	length := len(a.vals)
	if index >= length || index < 0 {
		return errs.NewErrIndexOutOfRange(length, index)
	}
	a.vals[index] = t
	return nil
}

func (a *ArrayList[T]) Len() int {
	return len(a.vals)
}

func (a *ArrayList[T]) Cap() int {
	return cap(a.vals)
}

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
	a.vals = slice.Shrink(a.vals)
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
