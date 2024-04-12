package mapx

type linkedKV[K any, V any] struct {
	key        K
	val        V
	prev, next *linkedKV[K, V]
}

type LinkedMap[K any, V any] struct {
	m          mapi[K, *linkedKV[K, V]]
	head, tail *linkedKV[K, V]
	length     int
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

// TODO
//func NewLinkedTreeMap[K any,V any]

func (lm *LinkedMap[K, V]) Put(key K, val V) error {
	if lk, ok := lm.m.Get(key); ok {
		lk.val = val
		return nil
	}
	lk := &linkedKV[K, V]{
		key:  key,
		val:  val,
		prev: lm.tail.prev,
		next: lm.tail,
	}
	if err := lm.m.Put(key, lk); err != nil {
		return err
	}
	lk.prev.next, lk.next.prev = lk, lk
	lm.length++
	return nil
}

func (lm *LinkedMap[K, V]) Get(key K) (V, bool) {
	if lk, ok := lm.m.Get(key); ok {
		return lk.val, ok
	}
	var v V
	return v, false
}

func (lm *LinkedMap[K, V]) Delete(key K) (V, bool) {
	if lk, ok := lm.m.Delete(key); ok {
		lk.prev.next = lk.next
		lk.next.prev = lk.prev
		lm.length--
		return lk.val, ok
	}
	var v V
	return v, false
}

func (lm *LinkedMap[K, V]) Keys() []K {
	keys := make([]K, 0, lm.length)
	for cur := lm.head.next; cur != lm.tail; {
		keys = append(keys, cur.key)
		cur = cur.next
	}
	return keys
}

func (lm *LinkedMap[K, V]) Values() []V {
	values := make([]V, 0, lm.length)
	for cur := lm.head.next; cur != lm.tail; {
		values = append(values, cur.val)
		cur = cur.next
	}
	return values
}

func (lm *LinkedMap[K, V]) Len() int64 {
	return int64(lm.length)
}
