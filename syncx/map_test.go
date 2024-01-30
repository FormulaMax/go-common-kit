package syncx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type User struct {
	Name string
}

func TestMap_Load(t *testing.T) {
	testCases := []struct {
		name    string
		key     string
		wantOk  bool
		wantVal *User
	}{
		{
			name:    "found",
			key:     "found",
			wantOk:  true,
			wantVal: &User{Name: "found"},
		},
		{
			name:    "found but empty",
			key:     "found but empty",
			wantOk:  true,
			wantVal: &User{},
		},
		{
			name: "not found",
			key:  "not found",
		},
	}

	var mu Map[string, *User]
	mu.Store("found", testCases[0].wantVal)
	mu.Store("found but empty", testCases[1].wantVal)

	mu.Store("found but nil", nil)
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			val, ok := mu.Load(tc.key)
			assert.Equal(t, tc.wantOk, ok)
			assert.Same(t, tc.wantVal, val)
		})
	}
}

func TestMap_LoadOrStore(t *testing.T) {
	t.Run("store non-nil value", func(t *testing.T) {
		m, user := Map[string, *User]{}, &User{Name: "Tom"}
		val, loaded := m.LoadOrStore(user.Name, user)
		assert.False(t, loaded)
		assert.Same(t, user, val)
	})
}
