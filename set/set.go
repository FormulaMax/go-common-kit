package set

type Set[T comparable] interface {
	Add(key T)
	Delete(key T)
	Exists(key T) bool
	Keys() []T
}

type MapSet[T comparable] struct {
	m map[T]struct{}
}

func NewMapSet[T comparable](size int) *MapSet[T] {
	return &MapSet[T]{
		m: make(map[T]struct{}, size),
	}
}

func (ms *MapSet[T]) Add(val T) {
	ms.m[val] = struct{}{}
}

func (ms *MapSet[T]) Delete(key T) {
	delete(ms.m, key)
}

func (ms *MapSet[T]) Exists(key T) bool {
	_, ok := ms.m[key]
	return ok
}

func (ms *MapSet[T]) Keys() []T {
	res := make([]T, 0, len(ms.m))
	for key := range ms.m {
		res = append(res, key)
	}
	return res
}
