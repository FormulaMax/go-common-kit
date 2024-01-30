package mapx

import (
	"errors"
	"github.com/FormulaMax/go-common-kit/internal/tree"
)

var (
	errTreeMapComparatorIsNull = errors.New("go-common:Comparator不能为nil")
)

type TreeMap[K any, V any] struct {
	tree *tree.RBTree[K, V]
}
