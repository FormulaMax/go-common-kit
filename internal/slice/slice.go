package slice

import "github.com/FormulaMax/go-common-kit/internal/errs"

func Add[T any](src []T, element T, index int) ([]T, error) {
	length := len(src)
	if index < 0 || index > length {
		return nil, errs.NewErrIndexOutOfRange(length, index)
	}
	var zeroVal T
	src = append(src, zeroVal)
	for i := len(src) - 1; i > index; i-- {
		if i-1 >= 0 {
			src[i] = src[i-1]
		}
	}
	src[index] = element
	return src, nil
}

func Delete[T any](src []T, index int) ([]T, T, error) {
	length := len(src)
	if index < 0 || index >= length {
		var t T
		return nil, t, errs.NewErrIndexOutOfRange(length, index)
	}
	res := src[index]
	for i := index; i+1 < len(src); i++ {
		src[i] = src[i+1]
	}
	src = src[:length-1]
	return src, res, nil
}

func calCapacity(cap, len int) (int, bool) {
	if cap <= 64 {
		return cap, false
	}
	if cap > 2048 && (cap/len >= 2) {
		factor := 0.625
		return int(float32(cap) * float32(factor)), true
	}
	if cap <= 2048 && (cap/len >= 4) {
		return cap / 2, true
	}
	return cap, false
}

func Shrink[T any](src []T) []T {
	cap, len := cap(src), len(src)
	n, changed := calCapacity(cap, len)
	if !changed {
		return src
	}
	s := make([]T, 0, n)
	s = append(s, src...)
	return s
}
