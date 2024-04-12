package list

import (
	"errors"
	go_common_kit "github.com/FormulaMax/go-common-kit"
	"github.com/FormulaMax/go-common-kit/internal/errs"
	"golang.org/x/exp/rand"
)

const (
	FactorP  = float32(0.25)
	MaxLevel = 32
)

type skipListNode[T any] struct {
	Val     T
	Forward []*skipListNode[T]
}

type SkipList[T any] struct {
	header  *skipListNode[T]
	level   int
	compare go_common_kit.Comparator[T]
	size    int
}

func newSkipListNode[T any](val T, level int) *skipListNode[T] {
	return &skipListNode[T]{
		Val:     val,
		Forward: make([]*skipListNode[T], level),
	}
}

func NewSkipList[T any](compare go_common_kit.Comparator[T]) *SkipList[T] {
	return &SkipList[T]{
		header: &skipListNode[T]{
			Forward: make([]*skipListNode[T], MaxLevel),
		},
		level:   1,
		compare: compare,
	}
}

func NewSkipListFromSlice[T any](slice []T, compare go_common_kit.Comparator[T]) *SkipList[T] {
	sl := NewSkipList[T](compare)
	for _, n := range slice {
		sl.Insert(n)
	}
	return sl
}

func (sl *SkipList[T]) AsSlice() []T {
	curr := sl.header
	slice := make([]T, 0, sl.size)
	for curr.Forward[0] != nil {
		slice = append(slice, curr.Forward[0].Val)
		curr = curr.Forward[0]
	}
	return slice
}

func (sl *SkipList[T]) randomLevel() int {
	level := 1
	p := FactorP
	for (rand.Int31() & 0xFFFF) < int32(p*0xFFFF) {
		level++
	}
	if level < MaxLevel {
		return level
	}
	return MaxLevel
}

func (sl *SkipList[T]) Search(target T) bool {
	curr, _ := sl.traverse(target, sl.level)
	curr = curr.Forward[0]
	return curr != nil && sl.compare(curr.Val, target) == 0
}

func (sl *SkipList[T]) traverse(val T, level int) (*skipListNode[T], []*skipListNode[T]) {
	update := make([]*skipListNode[T], MaxLevel)
	curr := sl.header
	for i := level - 1; i >= 0; i-- {
		for curr.Forward[i] != nil && sl.compare(curr.Forward[i].Val, val) < 0 {
			curr = curr.Forward[i]
		}
		update[i] = curr
	}
	return curr, update
}

func (sl *SkipList[T]) Insert(val T) {
	_, update := sl.traverse(val, sl.level)
	level := sl.randomLevel()
	if level > sl.level {
		for i := sl.level; i < level; i++ {
			update[i] = sl.header
		}
		sl.level = level
	}

	// 插入新节点
	newNode := newSkipListNode[T](val, level)
	for i := 0; i < level; i++ {
		newNode.Forward[i] = update[i].Forward[i]
		update[i].Forward[i] = newNode
	}
	sl.size += 1
}

func (sl *SkipList[T]) Len() int {
	return sl.size
}

func (sl *SkipList[T]) DeleteElement(target T) bool {
	curr, update := sl.traverse(target, sl.level)
	node := curr.Forward[0]
	if node == nil || sl.compare(node.Val, target) != 0 {
		return true
	}
	for i := 0; i < sl.level && update[i].Forward[i] == node; i++ {
		update[i].Forward[i] = node.Forward[i]
	}
	for sl.level > 1 && sl.header.Forward[sl.level-1] == nil {
		sl.level--
	}
	sl.level -= 1
	return true
}

func (sl *SkipList[T]) Peek() (T, error) {
	curr := sl.header
	curr = curr.Forward[0]
	var zero T
	if curr == nil {
		return zero, errors.New("跳表为空")
	}
	return curr.Val, nil
}

func (sl *SkipList[T]) Get(index int) (T, error) {
	var zero T
	if index < 0 || index >= sl.size {
		return zero, errs.NewErrIndexOutOfRange(sl.size, index)
	}
	curr := sl.header
	for i := 0; i <= index; i++ {
		curr = curr.Forward[0]
	}
	return curr.Val, nil
}
