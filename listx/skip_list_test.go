package listx

import (
	go_common_kit "github.com/FormulaMax/go-common-kit"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSkipList(t *testing.T) {
	testCases := []struct {
		name      string
		compare   go_common_kit.Comparator[int]
		wantSlice []int
	}{
		{
			name:      "new skip list",
			compare:   go_common_kit.ComparatorRealNumber[int],
			wantSlice: []int{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sl := NewSkipList(tc.compare)
			assert.Equal(t, tc.wantSlice, sl.AsSlice())
		})
	}
}

func TestSkipList_AsSlice(t *testing.T) {
	testCases := []struct {
		name      string
		compare   go_common_kit.Comparator[int]
		wantSlice []int
	}{
		{
			name:      "no err is ok",
			compare:   go_common_kit.ComparatorRealNumber[int],
			wantSlice: []int{},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sl := NewSkipList[int](tc.compare)
			assert.Equal(t, tc.wantSlice, sl.AsSlice())
		})
	}
}

func TestSkipList_Cap(t *testing.T) {
	testCases := []struct {
		name     string
		compare  go_common_kit.Comparator[int]
		wantSize int
	}{
		{
			name:     "no err is ok",
			compare:  go_common_kit.ComparatorRealNumber[int],
			wantSize: 0,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sl := NewSkipList[int](tc.compare)
			assert.Equal(t, tc.wantSize, sl.Cap())
		})
	}
}

func TestSkipList_DeleteElement(t *testing.T) {
	testCases := []struct {
		name     string
		compare  go_common_kit.Comparator[int]
		value    int
		wantBool bool
	}{
		{
			name:     "no err is ok",
			compare:  go_common_kit.ComparatorRealNumber[int],
			value:    1,
			wantBool: true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sl := NewSkipList[int](tc.compare)
			ok := sl.DeleteElement(tc.value)
			assert.Equal(t, tc.wantBool, ok)
		})
	}
}

func TestSkipList_Insert(t *testing.T) {
	testCases := []struct {
		name      string
		compare   go_common_kit.Comparator[int]
		key       int
		wantSlice []int
	}{
		{
			name:      "no err is ok",
			compare:   go_common_kit.ComparatorRealNumber[int],
			key:       1,
			wantSlice: []int{1},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sl := NewSkipList[int](tc.compare)
			sl.Insert(tc.key)
			assert.Equal(t, tc.wantSlice, sl.AsSlice())
		})
	}
}

func TestSkipList_Len(t *testing.T) {
	testCases := []struct {
		name     string
		compare  go_common_kit.Comparator[int]
		wantSize int
	}{
		{
			name:     "no err is ok",
			compare:  go_common_kit.ComparatorRealNumber[int],
			wantSize: 0,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sl := NewSkipList[int](tc.compare)
			assert.Equal(t, tc.wantSize, sl.Len())
		})
	}
}

func TestSkipList_Search(t *testing.T) {
	testCases := []struct {
		name     string
		compare  go_common_kit.Comparator[int]
		value    int
		wantBool bool
	}{
		{
			name:     "no err is ok",
			compare:  go_common_kit.ComparatorRealNumber[int],
			value:    1,
			wantBool: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sl := NewSkipList[int](tc.compare)
			ok := sl.Search(tc.value)
			assert.Equal(t, tc.wantBool, ok)
		})
	}
}
