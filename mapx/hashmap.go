package mapx

import "github.com/FormulaMax/go-common-kit/syncx"

type Hashable interface {

	// Code 返回该元素的哈希值
	// 注意：哈希值应该尽可能的均匀以避免冲突
	Code() uint64

	// Equals 比较两个元素是否相等。如果返回 true，那么我们会认为两个键是一样的。
	Equals(key any) bool
}

type node[K Hashable, V any] struct {
	key   K
	value V
	next  *node[K, V]
}

var _ mapi[Hashable, any] = (*HashMap[Hashable, any])(nil)

type HashMap[K Hashable, V any] struct {
	hashmap  map[uint64]*node[K, V]
	nodePool *syncx.Pool[*node[K, V]]
}

func (hm *HashMap[K, V]) newNode(key K, val V) *node[K, V] {
	newNode := hm.nodePool.Get()
	newNode.value = val
	newNode.key = key
	return newNode
}

func NewHashMap[K Hashable, V any](size int) *HashMap[K, V] {
	return &HashMap[K, V]{
		nodePool: syncx.NewPool[*node[K, V]](func() *node[K, V] {
			return &node[K, V]{}
		}),
		hashmap: make(map[uint64]*node[K, V], size),
	}
}

func (hm *HashMap[K, V]) Put(key K, val V) error {
	hash := key.Code()
	root, ok := hm.hashmap[hash]
	if !ok {
		hash = key.Code()
		newNode := hm.newNode(key, val)
		hm.hashmap[hash] = newNode
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
	newNode := hm.newNode(key, val)
	pre.next = newNode
	return nil
}

func (hm *HashMap[K, V]) Get(key K) (V, bool) {
	hash := key.Code()
	root, ok := hm.hashmap[hash]
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
func (hm *HashMap[K, V]) Delete(key K) (V, bool) {
	root, ok := hm.hashmap[key.Code()]
	if !ok {
		var val V
		return val, false
	}
	pre := root
	num := 0
	for root != nil {
		if root.key.Equals(key) {
			if num == 0 && root.next == nil {
				delete(hm.hashmap, key.Code())
			} else if num == 0 && root.next != nil {
				hm.hashmap[key.Code()] = root.next
			} else {
				pre.next = root.next
			}
			val := root.value
			root.formatting()
			hm.nodePool.Put(root)
			return val, true
		}
		num++
		pre = root
		root = root.next
	}
	var val V
	return val, false
}

// Keys 返回 Hashmap 里面的所有的 key。
// 注意：key 的顺序是随机的。
func (hm *HashMap[K, V]) Keys() []K {
	res := make([]K, 0)
	for _, bucketNode := range hm.hashmap {
		curNode := bucketNode
		for curNode != nil {
			res = append(res, curNode.key)
			curNode = curNode.next
		}
	}
	return res
}

// Values 返回 Hashmap 里面的所有的 value。
// 注意：value 的顺序是随机的。
func (hm *HashMap[K, V]) Values() []V {
	res := make([]V, 0)
	for _, bucketNode := range hm.hashmap {
		curNode := bucketNode
		for curNode != nil {
			res = append(res, curNode.value)
			curNode = curNode.next
		}
	}
	return res
}

func (hm *HashMap[K, V]) Len() int64 {
	return int64(len(hm.hashmap))
}

func (n *node[K, V]) formatting() {
	var val V
	var key K
	n.key = key
	n.value = val
	n.next = nil
}
