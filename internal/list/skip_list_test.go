package list

import (
	"errors"
	"fmt"
	go_common_kit "github.com/FormulaMax/go-common-kit"
	"github.com/FormulaMax/go-common-kit/internal/errs"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSkipList(t *testing.T) {
	testCases := []struct {
		name       string
		compare    go_common_kit.Comparator[int]
		level      int
		wantHeader *skipListNode[int]
		wantLevel  int
		wantSlice  []int
		wantErr    error
		wantSize   int
	}{
		{
			name:       "new skip list",
			compare:    go_common_kit.ComparatorRealNumber[int],
			level:      1,
			wantLevel:  1,
			wantHeader: newSkipListNode[int](0, MaxLevel),
			wantSlice:  []int{},
			wantSize:   0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sl := NewSkipList[int](tc.compare)
			assert.Equal(t, tc.wantLevel, sl.level)
			assert.Equal(t, tc.wantHeader, sl.header)
			assert.Equal(t, tc.wantSlice, sl.AsSlice())
			assert.Equal(t, tc.wantSize, sl.size)

		})
	}
}

func TestNewSkipListFromSlice(t *testing.T) {
	testCases := []struct {
		name    string
		compare go_common_kit.Comparator[int]
		level   int
		slice   []int

		wantSlice []int
		wantErr   error
		wantSize  int
	}{
		{
			name:    "new skip list",
			compare: go_common_kit.ComparatorRealNumber[int],
			level:   1,
			slice:   []int{1, 2, 3},

			wantSlice: []int{1, 2, 3},
			wantSize:  3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sl := NewSkipListFromSlice[int](tc.slice, tc.compare)
			assert.Equal(t, tc.wantSlice, sl.AsSlice())
			assert.Equal(t, tc.wantSize, sl.size)

		})
	}
}

func TestSkipList_DeleteElement(t *testing.T) {
	testCases := []struct {
		name      string
		skiplist  *SkipList[int]
		compare   go_common_kit.Comparator[int]
		value     int
		wantSlice []int
		wantSize  int
		wantRes   bool
	}{
		{
			name:      "delete 2 from [1,3]",
			compare:   go_common_kit.ComparatorRealNumber[int],
			skiplist:  NewSkipListFromSlice[int]([]int{1, 3}, go_common_kit.ComparatorRealNumber[int]),
			value:     2,
			wantSlice: []int{1, 3},
			wantSize:  2,
			wantRes:   true,
		},
		{
			name:      "delete 1 from [1,3]",
			compare:   go_common_kit.ComparatorRealNumber[int],
			skiplist:  NewSkipListFromSlice[int]([]int{1, 3}, go_common_kit.ComparatorRealNumber[int]),
			value:     1,
			wantSlice: []int{3},
			wantSize:  1,
			wantRes:   true,
		},
		{
			name:      "delete 1 from []",
			compare:   go_common_kit.ComparatorRealNumber[int],
			skiplist:  NewSkipListFromSlice[int]([]int{}, go_common_kit.ComparatorRealNumber[int]),
			value:     1,
			wantSlice: []int{},
			wantSize:  0,
			wantRes:   true,
		},
		{
			name:      "delete 1 from [1]",
			compare:   go_common_kit.ComparatorRealNumber[int],
			skiplist:  NewSkipListFromSlice[int]([]int{1}, go_common_kit.ComparatorRealNumber[int]),
			value:     1,
			wantSlice: []int{},
			wantSize:  0,
			wantRes:   true,
		},
		{
			name:      "delete 1 from [2]",
			compare:   go_common_kit.ComparatorRealNumber[int],
			skiplist:  NewSkipListFromSlice[int]([]int{2}, go_common_kit.ComparatorRealNumber[int]),
			value:     1,
			wantSlice: []int{2},
			wantSize:  1,
			wantRes:   true,
		},
		{
			name:      "delete 3 from [1,2,3,4,5,6,7]",
			compare:   go_common_kit.ComparatorRealNumber[int],
			skiplist:  NewSkipListFromSlice[int]([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}, go_common_kit.ComparatorRealNumber[int]),
			value:     3,
			wantSlice: []int{1, 2, 4, 5, 6, 7},
			wantSize:  6,
			wantRes:   true,
		},
		{
			name:      "delete 8 from [1,2,3,4,5,6,7]",
			compare:   go_common_kit.ComparatorRealNumber[int],
			skiplist:  NewSkipListFromSlice[int]([]int{1, 2, 3, 4, 5, 6, 7}, go_common_kit.ComparatorRealNumber[int]),
			value:     8,
			wantSlice: []int{1, 2, 3, 4, 5, 6, 7},
			wantSize:  7,
			wantRes:   true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ok := tc.skiplist.DeleteElement(tc.value)
			assert.Equal(t, tc.wantSize, tc.skiplist.size)
			assert.Equal(t, tc.wantSlice, tc.skiplist.AsSlice())
			assert.Equal(t, tc.wantRes, ok)
		})
	}
}

func TestSkipList_Insert(t *testing.T) {
	testCases := []struct {
		name      string
		skiplist  *SkipList[int]
		compare   go_common_kit.Comparator[int]
		value     int
		wantSlice []int
		wantSize  int
	}{
		{
			name:      "insert 2 into [1,3]",
			compare:   go_common_kit.ComparatorRealNumber[int],
			skiplist:  NewSkipListFromSlice[int]([]int{1, 3}, go_common_kit.ComparatorRealNumber[int]),
			value:     2,
			wantSlice: []int{1, 2, 3},
			wantSize:  3,
		},
		{
			name:      "insert 1 into []",
			compare:   go_common_kit.ComparatorRealNumber[int],
			skiplist:  NewSkipListFromSlice[int]([]int{}, go_common_kit.ComparatorRealNumber[int]),
			value:     1,
			wantSlice: []int{1},
			wantSize:  1,
		},
		{
			name:      "insert 2 into [1,2,3]",
			compare:   go_common_kit.ComparatorRealNumber[int],
			skiplist:  NewSkipListFromSlice[int]([]int{1, 2, 3}, go_common_kit.ComparatorRealNumber[int]),
			value:     2,
			wantSlice: []int{1, 2, 2, 3},
			wantSize:  4,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.skiplist.Insert(tc.value)
			assert.Equal(t, tc.wantSize, tc.skiplist.size)
			assert.Equal(t, tc.wantSlice, tc.skiplist.AsSlice())
		})
	}
}

func TestSkipList_Search(t *testing.T) {
	testCases := []struct {
		name      string
		skiplist  *SkipList[int]
		compare   go_common_kit.Comparator[int]
		value     int
		wantSlice []int
		wantSize  int
		wantRes   bool
	}{
		{
			name:      "search 2 from [1,3]",
			compare:   go_common_kit.ComparatorRealNumber[int],
			skiplist:  NewSkipListFromSlice[int]([]int{1, 3}, go_common_kit.ComparatorRealNumber[int]),
			value:     2,
			wantSlice: []int{1, 3},
			wantSize:  2,
			wantRes:   false,
		},
		{
			name:      "search 1 from [1,3]",
			compare:   go_common_kit.ComparatorRealNumber[int],
			skiplist:  NewSkipListFromSlice[int]([]int{1, 3}, go_common_kit.ComparatorRealNumber[int]),
			value:     1,
			wantSlice: []int{1, 3},
			wantSize:  2,
			wantRes:   true,
		},
		{
			name:      "search 1 from []",
			compare:   go_common_kit.ComparatorRealNumber[int],
			skiplist:  NewSkipListFromSlice[int]([]int{}, go_common_kit.ComparatorRealNumber[int]),
			value:     1,
			wantSlice: []int{},
			wantSize:  0,
			wantRes:   false,
		},
		{
			name:      "search 1 from [1]",
			compare:   go_common_kit.ComparatorRealNumber[int],
			skiplist:  NewSkipListFromSlice[int]([]int{1}, go_common_kit.ComparatorRealNumber[int]),
			value:     1,
			wantSlice: []int{1},
			wantSize:  1,
			wantRes:   true,
		},
		{
			name:      "search 1 from [2]",
			compare:   go_common_kit.ComparatorRealNumber[int],
			skiplist:  NewSkipListFromSlice[int]([]int{2}, go_common_kit.ComparatorRealNumber[int]),
			value:     1,
			wantSlice: []int{2},
			wantSize:  1,
			wantRes:   false,
		},
		{
			name:      "search 3 from [1,2,3,4,5,6]",
			compare:   go_common_kit.ComparatorRealNumber[int],
			skiplist:  NewSkipListFromSlice[int]([]int{1, 2, 3, 4, 5, 6}, go_common_kit.ComparatorRealNumber[int]),
			value:     3,
			wantSlice: []int{1, 2, 3, 4, 5, 6},
			wantSize:  6,
			wantRes:   true,
		},
		{
			name:      "search 8 from [1,2,3,4,5,6]",
			compare:   go_common_kit.ComparatorRealNumber[int],
			skiplist:  NewSkipListFromSlice[int]([]int{1, 2, 3, 4, 5, 6}, go_common_kit.ComparatorRealNumber[int]),
			value:     8,
			wantSlice: []int{1, 2, 3, 4, 5, 6},
			wantSize:  6,
			wantRes:   false,
		},
		{
			name:      "search 2 from [1,2,2,3,3,4,5,6]",
			compare:   go_common_kit.ComparatorRealNumber[int],
			skiplist:  NewSkipListFromSlice[int]([]int{1, 2, 2, 3, 3, 4, 5, 6}, go_common_kit.ComparatorRealNumber[int]),
			value:     2,
			wantSlice: []int{1, 2, 2, 3, 3, 4, 5, 6},
			wantSize:  8,
			wantRes:   true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ok := tc.skiplist.Search(tc.value)
			assert.Equal(t, tc.wantSize, tc.skiplist.size)
			assert.Equal(t, tc.wantSlice, tc.skiplist.AsSlice())
			assert.Equal(t, tc.wantRes, ok)
		})
	}
}

func TestSkipList_randomLevel(t *testing.T) {
	sl := NewSkipListFromSlice[int]([]int{1, 2, 3}, go_common_kit.ComparatorRealNumber[int])
	fmt.Println(sl.randomLevel())
}

func TestSkipList_Peek(t *testing.T) {
	testCases := []struct {
		name      string
		skiplist  *SkipList[int]
		compare   go_common_kit.Comparator[int]
		wantSlice []int
		wantVal   int
		wantErr   error
	}{
		{
			name:      "peek [1,3]",
			compare:   go_common_kit.ComparatorRealNumber[int],
			skiplist:  NewSkipListFromSlice[int]([]int{1, 3}, go_common_kit.ComparatorRealNumber[int]),
			wantSlice: []int{1, 3},
			wantVal:   1,
			wantErr:   nil,
		},
		{
			name:      "peek []",
			compare:   go_common_kit.ComparatorRealNumber[int],
			skiplist:  NewSkipListFromSlice[int]([]int{}, go_common_kit.ComparatorRealNumber[int]),
			wantSlice: []int{},
			wantVal:   0,
			wantErr:   errors.New("跳表为空"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			val, err := tc.skiplist.Peek()
			assert.Equal(t, tc.wantErr, err)
			assert.Equal(t, tc.wantVal, val)
		})
	}
}

func TestSkipList_Get(t *testing.T) {
	testCases := []struct {
		name      string
		skiplist  *SkipList[int]
		compare   go_common_kit.Comparator[int]
		index     int
		wantSlice []int
		wantVal   int
		wantErr   error
	}{
		{
			name:      "get index -1 [1, 2, 3]",
			compare:   go_common_kit.ComparatorRealNumber[int],
			skiplist:  NewSkipListFromSlice[int]([]int{1, 2, 3}, go_common_kit.ComparatorRealNumber[int]),
			index:     -1,
			wantSlice: []int{1, 2, 3},
			wantVal:   0,
			wantErr:   errs.NewErrIndexOutOfRange(3, -1),
		},
		{
			name:      "get index 3 [1, 2, 3]",
			compare:   go_common_kit.ComparatorRealNumber[int],
			skiplist:  NewSkipListFromSlice[int]([]int{1, 2, 3}, go_common_kit.ComparatorRealNumber[int]),
			index:     3,
			wantSlice: []int{1, 2, 3},
			wantVal:   0,
			wantErr:   errs.NewErrIndexOutOfRange(3, 3),
		},
		{
			name:      "get index 0 [1, 2, 3]",
			compare:   go_common_kit.ComparatorRealNumber[int],
			skiplist:  NewSkipListFromSlice[int]([]int{1, 2, 3}, go_common_kit.ComparatorRealNumber[int]),
			index:     0,
			wantSlice: []int{1, 2, 3},
			wantVal:   1,
			wantErr:   nil,
		},
		{
			name:      "get index 1 [1, 2, 3]",
			compare:   go_common_kit.ComparatorRealNumber[int],
			skiplist:  NewSkipListFromSlice[int]([]int{1, 2, 3}, go_common_kit.ComparatorRealNumber[int]),
			index:     1,
			wantSlice: []int{1, 2, 3},
			wantVal:   2,
			wantErr:   nil,
		},
		{
			name:      "get index 2 [1, 2, 3]",
			compare:   go_common_kit.ComparatorRealNumber[int],
			skiplist:  NewSkipListFromSlice[int]([]int{1, 2, 3}, go_common_kit.ComparatorRealNumber[int]),
			index:     2,
			wantSlice: []int{1, 2, 3},
			wantVal:   3,
			wantErr:   nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			val, err := tc.skiplist.Get(tc.index)
			assert.Equal(t, tc.wantErr, err)
			assert.Equal(t, tc.wantVal, val)
		})
	}
}

func TestSkipList_AsSlice(t *testing.T) {
	testCases := []struct {
		name      string
		skiplist  *SkipList[int]
		compare   go_common_kit.Comparator[int]
		wantSlice []int
	}{
		{
			name:      " [1, 2, 3]",
			compare:   go_common_kit.ComparatorRealNumber[int],
			skiplist:  NewSkipListFromSlice[int]([]int{1, 2, 3}, go_common_kit.ComparatorRealNumber[int]),
			wantSlice: []int{1, 2, 3},
		},
		{
			name:      "[3,2,1]]",
			compare:   go_common_kit.ComparatorRealNumber[int],
			skiplist:  NewSkipListFromSlice[int]([]int{3, 2, 1}, go_common_kit.ComparatorRealNumber[int]),
			wantSlice: []int{1, 2, 3},
		},
		{
			name:      "[]",
			compare:   go_common_kit.ComparatorRealNumber[int],
			skiplist:  NewSkipListFromSlice[int]([]int{1, 2, 3}, go_common_kit.ComparatorRealNumber[int]),
			wantSlice: []int{1, 2, 3},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.wantSlice, tc.skiplist.AsSlice())
		})
	}
}
