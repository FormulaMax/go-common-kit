package slice

import (
	"fmt"
)

// DeleteByIndex 指定下标删除
func DeleteByIndex[T any](src []T, index int) ([]T, T, error) {

	var target T

	//判定下标合法
	srcLength := len(src)
	if index < 0 || index > srcLength {
		return nil, target, fmt.Errorf("下标超界")
	}

	target = src[index]

	var res = make([]T, 0, cap(src))

	//复制index之前的元素
	for i := 0; i < index; i++ {
		res = append(res, src[i])
	}

	//复制index之后的元素
	for i := index + 1; i < len(src); i++ {
		res = append(res, src[i])
	}

	//检测是否需要缩容
	if isNeedToShrink(res) {
		res = shrinkSlice(res)
	}

	return res, target, nil
}

// DeleteByFilter 指定条件值删除
func DeleteByFilter[T any](src []T, method func(src T) bool) []T {

	var empty int = 0

	//cap := cap(src)
	length := len(src)
	for i := 0; i+empty < length; i++ {
		if method(src[i+empty]) {
			src[i] = src[empty+1]
			empty++
		}
	}

	//检测是否需要缩容
	if isNeedToShrink(src) {
		src = shrinkSlice(src)
	}

	return src
}

// isNeedToShrink 判定是否需要缩容
func isNeedToShrink[T any](src []T) bool {
	cap := cap(src)

	//容量小于256时再缩容
	if cap <= 256 {
		return false
	}

	if cap > 0 && len(src) < (cap>>1) {
		return true
	}

	return false
}

// shrinkSlice 缩容
func shrinkSlice[T any](src []T) []T {

	length := len(src)

	//容量减半
	var res = make([]T, 0, cap(src)/2)
	for i := length; i < length; i++ {
		res = append(res, src[i])
	}
	return res
}
