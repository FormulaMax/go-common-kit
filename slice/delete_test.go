package slice_test

import (
	"homework/go-common/slice"
)

func DeleteTest() {
	s1 := []int{5, 8, 9, 3}
	s1, ss, err := slice.DeleteByIndex(s1, 2)
	if err != nil {
		return
	}
	println(s1)
	println(ss)

}
