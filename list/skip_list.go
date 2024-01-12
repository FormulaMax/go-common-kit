package list

import (
	go_common_kit "github.com/FormulaMax/go-common-kit"
	"github.com/FormulaMax/go-common-kit/internal/list"
)

func NewSkipList[T any](compare go_common_kit.Comparator[T]) *SkipList[T] {
	pq := &SkipList[T]{}
	pq.skiplist = list.NewSkipList[T](compare)
	return pq
}

type SkipList[T any] struct {
	skiplist *list.SkipList[T]
}

func (sl *SkipList[T]) Search(target T) bool {
	return sl.skiplist.Search(target)
}

func (sl *SkipList[T]) AsSlice() []T {
	return sl.skiplist.AsSlice()
}

func (sl *SkipList[T]) Len() int {
	return sl.skiplist.Len()
}

func (sl *SkipList[T]) Cap() int {
	return sl.Len()
}

func (sl *SkipList[T]) Insert(Val T) {
	sl.skiplist.Insert(Val)
}

func (sl *SkipList[T]) DeleteElement(target T) bool {
	return sl.skiplist.DeleteElement(target)
}
