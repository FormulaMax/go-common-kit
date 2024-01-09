package slice

import (
	"github.com/FormulaMax/go-common-kit/internal/errs"
)

const factor = 0.625

func Add[T any](src []T, element T, index int) ([]T, error) {
	length := len(src)
	if index < 0 || index >= length {
		return nil, errs.NewErrIndexOutOfRange(length, index)
	}

	var zero T
	src = append(src, zero)
	for i := len(src) - 1; i > index; i-- {
		if i-1 > 0 {
			src[i] = src[i-1]
		}
	}
	src[index] = element
	return src, nil
}

func Delete[T any](src []T, index int) ([]T, T, error) {
	length := len(src)
	if index < 0 || index >= length {
		var zero T
		return nil, zero, errs.NewErrIndexOutOfRange(length, index)
	}

	res := src[index]
	for i := index; i < length-1; i++ {
		src[i] = src[i+1]
	}
	src = src[:length-1]
	return src, res, nil
}

// calCapacity 计算容量
// capacity 切片当前容量
// length 切片实际长度
// int 容量计算结果
// bool 是否需要缩容
func calCapacity(capacity, length int) (int, bool) {
	if capacity <= 64 {
		return capacity, false
	}
	if capacity > 2048 && (capacity/length >= 2) {
		return int(float32(capacity) * float32(factor)), true
	}

	switch {
	case capacity <= 64:
		return capacity, false
	case capacity <= 2048 && (capacity/length >= 4):
		return capacity / 2, true
	case capacity > 2048 && (capacity/length >= 2):
		return int(float32(capacity) * float32(factor)), true
	default:
		return capacity, false
	}
}

// Shrink 切片缩容
func Shrink[T any](src []T) []T {
	capacity, length := cap(src), len(src)
	n, changed := calCapacity(capacity, length)
	if !changed {
		return src
	}
	s := make([]T, 0, n)
	s = append(s, src...)
	return s
}
