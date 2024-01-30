package mapx

type MultiMap[K any, V any] struct {
	m mapx[K, []V]
}

//func NewMultiBuiltinMap[K comparable,V any](size int) *MultiMap[K,V]{
//	var m mapx[K,[]V]=newBuiltinMap(K,[]V)(size)
//	return &MultiMap[K,V]{
//		m: m,
//	}
//}
//
//func NewMultiHashMap[K comparable,V any](size int)*MultiMap[K,V]{
//	var m mapx[K,[]V]=NewHashMap[K,[]V](size)
//	return &MultiMap[K,V]{
//		m: m,
//	}
//}
//
//func NewMultiTreeMap[K any, V any](comparator go_common_kit.Comparator[K]) (*MultiMap[K, V], error) {
//	treeMap, err := NewTreeMap[K, []V](comparator)
//	if err != nil {
//		return nil, err
//	}
//	return &MultiMap[K, V]{
//		m: treeMap,
//	}, nil
//}
