package multiset

import (
	"cmp"
	"time"
)

/*
对标 C++ 的 map
没有重复 key
*/

type node[K comparable, V any] struct {
	son      [2]*node[K, V]
	priority uint
	key      K
	value    V
	subSize  int
}

func (o *node[K, V]) size() int {
	if o != nil {
		return o.subSize
	}
	return 0
}

func (o *node[K, V]) maintain() {
	o.subSize = 1 + o.son[0].size() + o.son[1].size()
}

func (o *node[K, V]) rotate(d int) *node[K, V] {
	x := o.son[d^1]
	o.son[d^1] = x.son[d]
	x.son[d] = o
	o.maintain()
	x.maintain()
	return x
}

type treap[K comparable, V any] struct {
	rd         uint
	root       *node[K, V]
	comparator func(a, b K) int
}

func (t *treap[K, V]) fastRand() uint {
	t.rd ^= t.rd << 13
	t.rd ^= t.rd >> 17
	t.rd ^= t.rd << 5
	return t.rd
}

func (t *treap[K, V]) size() int   { return t.root.size() }
func (t *treap[K, V]) empty() bool { return t.size() == 0 }

func (t *treap[K, V]) _put(o *node[K, V], key K, value V) *node[K, V] {
	if o == nil {
		o = &node[K, V]{priority: t.fastRand(), key: key, value: value}
	} else {
		c := t.comparator(key, o.key)
		if c == 0 { // 相等
			o.value = value
		} else {
			// c < 0 去左边，c > 0 去右边
			d := 0
			if c > 0 {
				d = 1
			}
			o.son[d] = t._put(o.son[d], key, value)
			if o.son[d].priority > o.priority {
				o = o.rotate(d ^ 1)
			}
		}
	}
	o.maintain()
	return o
}

func (t *treap[K, V]) put(key K, value V) { t.root = t._put(t.root, key, value) }

func (t *treap[K, V]) _delete(o *node[K, V], key K) *node[K, V] {
	if o == nil {
		return nil
	}
	if c := t.comparator(key, o.key); c != 0 {
		// < 0 去左边，> 0 去右边
		d := 0
		if c > 0 {
			d = 1
		}
		o.son[d] = t._delete(o.son[d], key)
	} else { // 相等，删除
		if o.son[1] == nil {
			return o.son[0]
		}
		if o.son[0] == nil {
			return o.son[1]
		}
		d := 0
		if o.son[0].priority > o.son[1].priority {
			d = 1
		}
		o = o.rotate(d)
		o.son[d] = t._delete(o.son[d], key)
	}
	o.maintain()
	return o
}

func (t *treap[K, V]) delete(key K) { t.root = t._delete(t.root, key) }

func (t *treap[K, V]) min() *node[K, V] { return t.kth(0) }
func (t *treap[K, V]) max() *node[K, V] { return t.kth(t.size() - 1) }

// 把 treap 当作一个有序数组（自动去重），返回第一个 >= key 的数的下标，若不存在，返回 size()
// 等价于 < key 的元素个数
func (t *treap[K, V]) lowerBoundIndex(key K) (kth int) {
	for o := t.root; o != nil; {
		c := t.comparator(key, o.key)
		if c < 0 {
			o = o.son[0]
		} else if c > 0 {
			kth += o.son[0].size() + 1
			o = o.son[1]
		} else { // 相等
			kth += o.son[0].size()
			break
		}
	}
	return
}

// 把 treap 当作一个有序数组（自动去重），返回第一个 > key 的数的下标，若不存在，返回 size()
// 等价于 <= key 的元素个数
func (t *treap[K, V]) upperBoundIndex(key K) (kth int) {
	for o := t.root; o != nil; {
		c := t.comparator(key, o.key)
		if c < 0 {
			o = o.son[0]
		} else if c > 0 {
			kth += o.son[0].size() + 1
			o = o.son[1]
		} else { // 相等
			kth += o.son[0].size() + 1
			break
		}
	}
	return
}

// 把 treap 当作一个有序数组（自动去重），返回下标为 k 的 node（k 从 0 开始）
// 也就是第 k+1 小：node.key 是最小的满足「有至少 k+1 个元素 <= node.key」的元素
func (t *treap[K, V]) kth(k int) (o *node[K, V]) {
	if k < 0 || k >= t.root.size() {
		return // NOTE: check nil
	}
	for o = t.root; o != nil; {
		leftSize := o.son[0].size()
		if k < leftSize {
			o = o.son[0]
		} else {
			k -= leftSize + 1
			if k < 0 {
				break
			}
			o = o.son[1]
		}
	}
	return // NOTE: check nil
}

// 求 v 的前驱（小于 key 的最大元素）
func (t *treap[K, V]) prev(key K) *node[K, V] { return t.kth(t.lowerBoundIndex(key) - 1) }

// 求 v 的后继（大于 key 的最小元素）
func (t *treap[K, V]) next(key K) *node[K, V] { return t.kth(t.upperBoundIndex(key)) }

// <= key 可以用 t.kth(t.upperBoundIndex(key)-1)
// >= key 可以用 t.kth(t.lowerBoundIndex(key))

func (t *treap[K, V]) find(key K) *node[K, V] {
	o := t.kth(t.lowerBoundIndex(key))
	if o == nil || o.key != key {
		return nil
	}
	return o
}

func newTreap[K cmp.Ordered, V any]() *treap[K, V] {
	return &treap[K, V]{
		rd:         uint(time.Now().UnixNano())/2 + 1,
		comparator: cmp.Compare[K],
	}
}

func newTreapWith[K comparable, V any](comp func(a, b K) int) *treap[K, V] {
	return &treap[K, V]{
		rd:         uint(time.Now().UnixNano())/2 + 1,
		comparator: comp,
	}
}

// 具体用法见本目录下的 test.go 文件
