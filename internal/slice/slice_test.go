package slice

import (
	"github.com/FormulaMax/go-common-kit/internal/errs"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd(t *testing.T) {
	testCases := []struct {
		name  string
		slice []int
		val   int
		index int

		wantErr   error
		wantSlice []int
	}{
		{
			name:      "index 0",
			slice:     []int{123, 100},
			val:       233,
			index:     0,
			wantSlice: []int{233, 123, 100},
		},
		{
			name:      "index middle",
			slice:     []int{123, 124, 125},
			val:       233,
			index:     1,
			wantSlice: []int{123, 233, 124, 125},
		},
		{
			name:      "index out of range",
			slice:     []int{123, 100},
			val:       233,
			index:     12,
			wantSlice: nil,
			wantErr:   errs.NewErrIndexOutOfRange(2, 12),
		},
		{
			name:      "index less than 0",
			slice:     []int{123, 100},
			val:       233,
			index:     -1,
			wantSlice: nil,
			wantErr:   errs.NewErrIndexOutOfRange(2, -1),
		},
		{
			name:      "index last",
			slice:     []int{123, 124, 125},
			val:       233,
			index:     2,
			wantSlice: []int{123, 124, 233, 125},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := Add[int](tc.slice, tc.val, tc.index)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantSlice, res)
		})
	}
}

func TestDelete(t *testing.T) {
	testCases := []struct {
		name      string
		slice     []int
		index     int
		wantErr   error
		wantSlice []int
		wantVal   int
	}{
		{
			name:      "empty slice",
			slice:     []int{},
			index:     0,
			wantErr:   errs.NewErrIndexOutOfRange(0, 0),
			wantSlice: nil,
			wantVal:   0,
		},
		{
			name:      "index out of range",
			slice:     []int{1},
			index:     1,
			wantErr:   errs.NewErrIndexOutOfRange(1, 1),
			wantSlice: nil,
		},
		{
			name:      "index 0",
			slice:     []int{123, 100},
			index:     0,
			wantSlice: []int{100},
			wantVal:   123,
		},
		{
			name:      "index middle",
			slice:     []int{123, 124, 125},
			index:     1,
			wantSlice: []int{123, 125},
			wantVal:   124,
		},
		{
			name:      "index less than 0",
			slice:     []int{123, 100},
			index:     -1,
			wantSlice: nil,
			wantErr:   errs.NewErrIndexOutOfRange(2, -1),
			wantVal:   0,
		},
		{
			name:      "index last",
			slice:     []int{123, 124, 125},
			index:     2,
			wantSlice: []int{123, 124},
			wantVal:   125,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, val, err := Delete[int](tc.slice, tc.index)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantSlice, res)
			assert.Equal(t, tc.wantVal, val)
		})
	}
}

func TestShrink(t *testing.T) {
	testCases := []struct {
		name        string
		originCap   int
		enqueueLoop int
		expectCap   int
	}{
		{
			name:        "less than 64",
			originCap:   32,
			enqueueLoop: 6,
			expectCap:   32,
		},
		{
			name:        "less than 2048, length less than 1/4",
			originCap:   1000,
			enqueueLoop: 20,
			expectCap:   500,
		},
		{
			name:        "less than 2048, length greater than 1/4",
			originCap:   1000,
			enqueueLoop: 400,
			expectCap:   1000,
		},
		{
			name:        "greater than 2048, length less than 1/2",
			originCap:   3000,
			enqueueLoop: 60,
			expectCap:   1875,
		},
		{
			name:        "greater than 2048, length greater than 1/2",
			originCap:   3000,
			enqueueLoop: 2000,
			expectCap:   3000,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			l := make([]int, 0, tc.originCap)
			for i := 0; i < tc.enqueueLoop; i++ {
				l = append(l, i)
			}
			l = Shrink[int](l)
			assert.Equal(t, tc.expectCap, cap(l))
		})
	}
}
