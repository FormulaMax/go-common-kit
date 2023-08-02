package main

import (
	"fmt"
	"homework/go-common/slice"
)

func main() {
	var s1 = []int{5, 8, 9, 6, 3, 0, 0, 0, 0, 0, 1}
	//var s1 = make([]int, 5, 11)
	//s1[0] = 5
	//s1[1] = 9
	s1, _, _ = slice.DeleteByIndex(s1, 4)
	fmt.Println(s1)
}
