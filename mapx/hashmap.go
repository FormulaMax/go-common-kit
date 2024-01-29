package mapx

import "github.com/FormulaMax/go-common-kit/syncx"

type Hashable interface {
	Code() uint64
	Equals(key any) bool
}

type node[K Hashable, V any] struct {
	key   K
	value V
	next  *node[K, V]
}

type HashMap[K Hashable, V any] struct {
	hashmap  map[uint64]*node[K, V]
	nodePool *syncx.Pool[*node[K, V]]
}

func (m *HashMap[K, V]) newNode(key K, val V) *node[K, V] {
	newNode := m.nodePool.Get()
	newNode.key = key
	newNode.value = val
	return newNode
}

var _ mapx[Hashable, any] = (*HashMap[Hashable, any])(nil)

func NewHashMap[K Hashable, V any](size int) *HashMap[K, V] {
	return &HashMap[K, V]{
		nodePool: syncx.NewPool[*node[K, V]](func() *node[K, V] {
			return &node[K, V]{}
		}),
		hashmap: make(map[uint64]*node[K, V], size),
	}
}

func (m *HashMap[K, V]) Put(key K, val V) error {
	hash := key.Code()
	root, ok := m.hashmap[hash]
	if !ok {
		hash = key.Code()
		newNode := m.newNode(key, val)
		m.hashmap[hash] = newNode
		return nil
	}
	pre := root
	for root != nil {
		if root.key.Equals(key) {
			root.value = val
			return nil
		}
		pre = root
		root = root.next
	}
	newNode := m.newNode(key, val)
	pre.next = newNode
	return nil
}

func (m *HashMap[K, V]) Get(key K) (V, bool) {
	hash := key.Code()
	root, ok := m.hashmap[hash]
	var val V
	if !ok {
		return val, false
	}
	for root != nil {
		if root.key.Equals(key) {
			return root.value, true
		}
		root = root.next
	}
	return val, false
}

// Delete 第一个返回值为删除key的值，第二个是hashmap是否真的有这个key
func (m *HashMap[K, V]) Delete(key K) (V, bool) {
	root, ok := m.hashmap[key.Code()]
	if !ok {
		var t V
		return t, false
	}
	pre := root
	num := 0
	for root != nil {
		if root.key.Equals(key) {
			if num == 0 && root.next == nil {
				delete(m.hashmap, key.Code())
			} else if num == 0 && root.next != nil {
				m.hashmap[key.Code()] = root.next
			} else {
				pre.next = root.next
			}
			val := root.value
			root.formatting()
			m.nodePool.Put(root)
			return val, true
		}
		num++
		pre = root
		root = root.next
	}
	var t V
	return t, false
}

func (m *HashMap[K, V]) Keys() []K {
	res := make([]K, 0)
	for _, bucketNode := range m.hashmap {
		curNode := bucketNode
		for curNode != nil {
			res = append(res, curNode.key)
			curNode = curNode.next
		}
	}
	return res
}

func (m *HashMap[K, V]) Values() []V {
	res := make([]V, 0)
	for _, bucketNode := range m.hashmap {
		curNode := bucketNode
		for curNode != nil {
			res = append(res, curNode.value)
			curNode = curNode.next
		}
	}
	return res
}

func (n *node[K, V]) formatting() {
	var val V
	var t K
	n.key = t
	n.value = val
	n.next = nil
}
