package slice

import (
	"fmt"
	"github.com/FormulaMax/go-common-kit/internal/errs"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd(t *testing.T) {
	// Add 主要依赖于 internal/slice.Add 来保证正确性
	testCases := []struct {
		name      string
		slice     []int
		addVal    int
		index     int
		wantSlice []int
		wantErr   error
	}{
		{
			name:      "index 0",
			slice:     []int{123, 100},
			addVal:    233,
			index:     0,
			wantSlice: []int{233, 123, 100},
		},
		{
			name:    "index -1",
			slice:   []int{123, 100},
			index:   -1,
			wantErr: errs.NewErrIndexOutOfRange(2, -1),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := Add(tc.slice, tc.addVal, tc.index)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantSlice, res)
		})
	}
}

func ExampleAdd() {
	res, _ := Add[int]([]int{1, 2, 3, 4}, 233, 2)
	fmt.Println(res)
	_, err := Add[int]([]int{1, 2, 3, 4}, 233, -1)
	fmt.Println(err)
	// Output:
	// [1 2 233 3 4]
	// go-common-kit: 下标超出范围，长度 4, 下标 -1
}

func TestDelete(t *testing.T) {
	// Delete 主要依赖于 internal/slice.Delete 来保证正确性
	testCases := []struct {
		name      string
		slice     []int
		index     int
		wantSlice []int
		wantErr   error
	}{
		{
			name:      "index 0",
			slice:     []int{123, 100},
			index:     0,
			wantSlice: []int{100},
		},
		{
			name:    "index -1",
			slice:   []int{123, 100},
			index:   -1,
			wantErr: errs.NewErrIndexOutOfRange(2, -1),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := Delete(tc.slice, tc.index)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantSlice, res)
		})
	}
}

func ExampleDelete() {
	res, _ := Delete[int]([]int{1, 2, 3, 4}, 2)
	fmt.Println(res)
	_, err := Delete[int]([]int{1, 2, 3, 4}, -1)
	fmt.Println(err)
	// Output:
	// [1 2 4]
	// go-common-kit: 下标超出范围，长度 4, 下标 -1
}

func TestFilterDelete(t *testing.T) {

	testCases := []struct {
		name            string
		src             []int
		deleteCondition func(idx int, src int) bool

		wantRes []int
	}{
		{
			name: "空切片",
			src:  []int{},
			deleteCondition: func(idx int, src int) bool {
				return false
			},

			wantRes: []int{},
		},
		{
			name: "不删除元素",
			src:  []int{0, 1, 2, 3, 4, 5, 6, 7},
			deleteCondition: func(idx int, src int) bool {
				return false
			},

			wantRes: []int{0, 1, 2, 3, 4, 5, 6, 7},
		},
		{
			name: "删除首位元素",
			src:  []int{0, 1, 2, 3, 4, 5, 6},
			deleteCondition: func(idx int, src int) bool {
				return idx == 0
			},

			wantRes: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name: "删除前面两个元素",
			src:  []int{0, 1, 2, 3, 4, 5, 6, 7},
			deleteCondition: func(idx int, src int) bool {
				return idx == 0 || idx == 1
			},

			wantRes: []int{2, 3, 4, 5, 6, 7},
		},
		{
			name: "删除中间单个元素",
			src:  []int{0, 1, 2, 3, 4, 5, 6, 7},
			deleteCondition: func(idx int, src int) bool {
				return idx == 3
			},

			wantRes: []int{0, 1, 2, 4, 5, 6, 7},
		},
		{
			name: "删除中间多个不连续元素",
			src:  []int{0, 1, 2, 3, 4, 5, 6, 7},
			deleteCondition: func(idx int, src int) bool {
				return idx == 2 || idx == 4
			},

			wantRes: []int{0, 1, 3, 5, 6, 7},
		},
		{
			name: "删除中间多个连续元素",
			src:  []int{0, 1, 2, 3, 4, 5, 6, 7},
			deleteCondition: func(idx int, src int) bool {
				return idx == 3 || idx == 4
			},

			wantRes: []int{0, 1, 2, 5, 6, 7},
		},
		{
			name: "删除中间多个元素，第一部分为一个元素，第二部分为连续元素",
			src:  []int{0, 1, 2, 3, 4, 5, 6, 7},
			deleteCondition: func(idx int, src int) bool {
				return idx == 2 || idx == 4 || idx == 5
			},

			wantRes: []int{0, 1, 3, 6, 7},
		},
		{
			name: "删除中间多个元素，第一部分为连续元素，第二部分为一个元素",
			src:  []int{0, 1, 2, 3, 4, 5, 6, 7},
			deleteCondition: func(idx int, src int) bool {
				return idx == 2 || idx == 3 || idx == 5
			},

			wantRes: []int{0, 1, 4, 6, 7},
		},
		{
			name: "删除后面两个元素",
			src:  []int{0, 1, 2, 3, 4, 5, 6, 7},
			deleteCondition: func(idx int, src int) bool {
				return idx == 6 || idx == 7
			},

			wantRes: []int{0, 1, 2, 3, 4, 5},
		},
		{
			name: "删除末尾元素",
			src:  []int{0, 1, 2, 3, 4, 5, 6, 7},
			deleteCondition: func(idx int, src int) bool {
				return idx == 7
			},

			wantRes: []int{0, 1, 2, 3, 4, 5, 6},
		},
		{
			name: "删除所有元素",
			src:  []int{0, 1, 2, 3, 4, 5, 6, 7},
			deleteCondition: func(idx int, src int) bool {
				return true
			},

			wantRes: []int{},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := FilterDelete(tc.src, tc.deleteCondition)
			assert.Equal(t, tc.wantRes, res)
		})
	}
}
