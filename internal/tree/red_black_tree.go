package tree

import (
	"errors"
	go_common_kit "github.com/FormulaMax/go-common-kit"
)

type color bool

const (
	Red   color = false
	Black color = true
)

var (
	ErrRBTreeSameRBNode = errors.New("go-common-kit: RBTree不能添加重复节点Key")
	ErrRBTreeNotRBNode  = errors.New("go-common-kit: RBTree不存在节点Key")
)

// TODO

type rbNode[K any, V any] struct {
	color               color
	key                 K
	value               V
	left, right, parent *rbNode[K, V]
}

type RBTree[K any, V any] struct {
	root    *rbNode[K, V]
	compare go_common_kit.Comparator[K]
}
