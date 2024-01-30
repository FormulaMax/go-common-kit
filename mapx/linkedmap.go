package mapx

type LinkedMap[K any, V any] struct {
	m      mapx[K, *linkedKV[K, V]]
	head   *linkedKV[K, V]
	tail   *linkedKV[K, V]
	length int
}

type linkedKV[K any, V any] struct {
	key  K
	val  V
	prev *linkedKV[K, V]
	next *linkedKV[K, V]
}

func NewLinkedHashMap[K Hashable, V any](size int) *LinkedMap[K, V] {
	hashmap := NewHashMap[K, *linkedKV[K, V]](size)
	head := &linkedKV[K, V]{}
	tail := &linkedKV[K, V]{next: head, prev: head}
	head.prev, head.next = tail, tail
	return &LinkedMap[K, V]{
		m:    hashmap,
		head: head,
		tail: tail,
	}
}

//func NewLinkedTreeMap[K any, V any](comparator go_common_kit.Comparator[K]) (*LinkedMap[K, V], error) {
//	treeMap, err := NewTreeMap
//}

func (l *LinkedMap[K, V]) Put(key K, val V) error {
	if lk, ok := l.m.Get(key); ok {
		lk.val = val
		return nil
	}
	lk := &linkedKV[K, V]{
		key:  key,
		val:  val,
		prev: l.tail.prev,
		next: l.tail,
	}
	if err := l.m.Put(key, lk); err != nil {
		return err
	}
	lk.prev.next, lk.next.prev = lk, lk
	l.length++
	return nil
}

func (l *LinkedMap[K, V]) Get(key K) (V, bool) {
	if lk, ok := l.m.Get(key); ok {
		lk.prev.next = lk.next
		lk.next.prev = lk.prev
		l.length--
		return lk.val, true
	}
	var val V
	return val, false
}

func (l *LinkedMap[K, V]) Keys() []K {
	keys := make([]K, 0, l.length)
	for cur := l.head.next; cur != l.tail; {
		keys = append(keys, cur.key)
		cur = cur.next
	}
	return keys
}

func (l *LinkedMap[K, V]) Values() []V {
	values := make([]V, 0, l.length)
	for cur := l.head.next; cur != l.tail; {
		values = append(values, cur.val)
		cur = cur.next
	}
	return values
}
