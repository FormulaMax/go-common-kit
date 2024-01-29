package mapx

type builtMap[K comparable, V any] struct {
	data map[K]V
}

func (b *builtMap[K, V]) Put(key K, val V) error {
	b.data[key] = val
	return nil
}

func (b *builtMap[K, V]) Get(key K) (V, bool) {
	val, ok := b.data[key]
	return val, ok
}

func (b *builtMap[K, V]) Delete(key K) (V, bool) {
	v, ok := b.data[key]
	delete(b.data, key)
	return v, ok
}

func (b *builtMap[K, V]) Keys() []K {
	return Keys[K, V](b.data)
}

func (b *builtMap[K, V]) Values() []V {
	return Values[K, V](b.data)
}

func newBuiltinMap[K comparable, V any](capacity int) *builtMap[K, V] {
	return &builtMap[K, V]{
		data: make(map[K]V, capacity),
	}
}
