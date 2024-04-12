package mapx

// builtinMap 是对 map 的二次封装
// 主要用于各种装饰器模式中被装饰的那个
type builtinMap[K comparable, V any] struct {
	data map[K]V
}

func (b *builtinMap[K, V]) Put(key K, val V) error {
	b.data[key] = val
	return nil
}

func (b *builtinMap[K, V]) Get(key K) (V, bool) {
	val, ok := b.data[key]
	return val, ok
}

func (b *builtinMap[K, V]) Delete(key K) (V, bool) {
	v, ok := b.data[key]
	delete(b.data, key)
	return v, ok
}

func (b *builtinMap[K, V]) Keys() []K {
	return Keys[K, V](b.data)
}

func (b *builtinMap[K, V]) Values() []V {
	return Values[K, V](b.data)
}

func newBuiltinMap[K comparable, V any](capacity int) *builtinMap[K, V] {
	return &builtinMap[K, V]{
		data: make(map[K]V, capacity),
	}
}

func (b *builtinMap[K, V]) Len() int64 {
	return int64(len(b.data))
}
