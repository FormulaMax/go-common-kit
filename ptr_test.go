package go_common_kit

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToPtr(t *testing.T) {
	i := 1285
	res := ToPtr[int](i)
	fmt.Println(&i)
	fmt.Println(res)
	assert.Equal(t, &i, res)
}
