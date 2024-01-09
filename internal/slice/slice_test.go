package slice

import (
	"github.com/FormulaMax/go-common-kit/internal/errs"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDelete(t *testing.T) {
	testCases := []struct {
		name  string
		slice []int
		index int

		wantErr   error
		wantSlice []int
	}{
		{
			name: "index out of range",
		},
		{
			name:      "empty slice",
			slice:     []int{},
			index:     0,
			wantErr:   errs.NewErrIndexOutOfRange(0, 0),
			wantSlice: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, _, err := Delete[int](tc.slice, tc.index)
			assert.Equal(t, tc.wantErr, err)
			assert.Equal(t, tc.wantSlice, res)
		})
	}
}
