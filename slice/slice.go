package slice

import "github.com/FormulaMax/go-common-kit/internal/slice"

// Add 在index处添加元素
// index 范围应为[0, len(src)]
// 如果index == len(src) 则表示往末尾添加元素
func Add[T any](src []T, element T, index int) ([]T, error) {
	res, err := slice.Add[T](src, element, index)
	return res, err
}

// Delete 删除 index 处的元素
func Delete[T any](src []T, index int) ([]T, error) {
	res, _, err := slice.Delete[T](src, index)
	return res, err
}

// FilterDelete 删除符合条件的元素
// 考虑到性能问题，所有操作都会在原切片上进行
// 被删除元素之后的元素会往前移动，有且只会移动一次
func FilterDelete[T any](src []T, m func(idx int, src T) bool) []T {
	// 记录被删除的元素位置，也称空缺的位置
	emptyPos := 0
	for idx := range src {
		// 判断是否满足删除的条件
		if m(idx, src[idx]) {
			continue
		}
		// 移动元素
		src[emptyPos] = src[idx]
		emptyPos++
	}
	return src[:emptyPos]
}
