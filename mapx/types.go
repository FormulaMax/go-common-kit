package mapx

type mapx[K any, V any] interface {
	Put(key K, val V) error
	Get(key K) (V, bool)
	Delete(key K) (V, bool)
	Keys() []K
	Values() []V
}
