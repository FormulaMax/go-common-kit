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
	ErrRBTreeSameRBNode = errors.New("go-common: RBTree不能添加重复节点Key")
	ErrRBTreeNotRBNode  = errors.New("go-common: RBTree不存在节点Key")
)

type rbNode[K any, V any] struct {
	color               color
	key                 K
	value               V
	left, right, parent *rbNode[K, V]
}

type RBTree[K any, V any] struct {
	root    *rbNode[K, V]
	compare go_common_kit.Comparator[K]
	size    int
}

func NewRBTree[K any, V any](compare go_common_kit.Comparator[K]) *RBTree[K, V] {
	return &RBTree[K, V]{
		compare: compare,
		root:    nil,
	}
}

func newRBNode[K any, V any](key K, value V) *rbNode[K, V] {
	return &rbNode[K, V]{
		key:    key,
		value:  value,
		color:  Red,
		left:   nil,
		right:  nil,
		parent: nil,
	}
}

func (rb *RBTree[K, V]) Size() int {
	if rb == nil {
		return 0
	}
	return rb.size
}

func (node *rbNode[K, V]) setNode(v V) {
	if node == nil {
		return
	}
	node.value = v
}

//func (rb *RBTree[K, V]) Add(key K, val V) error {
//	return rb.addNode(newRBNode(key, val))
//}
