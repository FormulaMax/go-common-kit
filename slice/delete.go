package slice

import "fmt"

// DeleteByIndex 指定下标删除
func DeleteByIndex[T any](src []T, index int) ([]T, T, error) {

	var defaultOfT T

	//判定下标合法
	srcLength := len(src)
	if index < 0 || index > srcLength {
		return nil, defaultOfT, fmt.Errorf("下标超界")
	}

	return nil, defaultOfT, nil
}

// DeleteByFilter 指定条件值删除
func DeleteByFilter[T any](src []any, val any, method func(idx int, src any) bool) []any {

	var targetIndex int = 0

	for index := range src {
		if method(index, src[index]) {
			continue
		}
		src[targetIndex] = src[index]

		targetIndex++
	}
	return src[:targetIndex]
}
