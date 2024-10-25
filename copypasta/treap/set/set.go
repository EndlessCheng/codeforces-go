package multiset

import (
	"cmp"
	"time"
)

/*
对标 C++ 的 set
没有重复 key
*/

type node[K comparable] struct {
	son      [2]*node[K]
	priority uint
	key      K
	subSize  int
}

func (o *node[K]) size() int {
	if o != nil {
		return o.subSize
	}
	return 0
}

func (o *node[K]) maintain() {
	o.subSize = 1 + o.son[0].size() + o.son[1].size()
}

func (o *node[K]) rotate(d int) *node[K] {
	x := o.son[d^1]
	o.son[d^1] = x.son[d]
	x.son[d] = o
	o.maintain()
	x.maintain()
	return x
}

type treap[K comparable] struct {
	rd         uint
	root       *node[K]
	comparator func(a, b K) int
}

func (t *treap[K]) fastRand() uint {
	t.rd ^= t.rd << 13
	t.rd ^= t.rd >> 17
	t.rd ^= t.rd << 5
	return t.rd
}

func (t *treap[K]) size() int   { return t.root.size() }
func (t *treap[K]) empty() bool { return t.size() == 0 }

func (t *treap[K]) _put(o *node[K], key K) *node[K] {
	if o == nil {
		o = &node[K]{priority: t.fastRand(), key: key}
	} else {
		c := t.comparator(key, o.key)
		if c == 0 {
			// 相等，什么也不做
		} else {
			// c < 0 去左边，c > 0 去右边
			d := 0
			if c > 0 {
				d = 1
			}
			o.son[d] = t._put(o.son[d], key)
			if o.son[d].priority > o.priority {
				o = o.rotate(d ^ 1)
			}
		}
	}
	o.maintain()
	return o
}

func (t *treap[K]) put(key K) { t.root = t._put(t.root, key) }

func (t *treap[K]) _delete(o *node[K], key K) *node[K] {
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

func (t *treap[K]) delete(key K) { t.root = t._delete(t.root, key) }

func (t *treap[K]) min() *node[K] { return t.kth(0) }
func (t *treap[K]) max() *node[K] { return t.kth(t.size() - 1) }

// 把 treap 当作一个有序数组（自动去重），返回第一个 >= key 的数的下标，若不存在，返回 size()
// 等价于 < key 的元素个数
func (t *treap[K]) lowerBoundIndex(key K) (kth int) {
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
func (t *treap[K]) upperBoundIndex(key K) (kth int) {
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
func (t *treap[K]) kth(k int) (o *node[K]) {
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
func (t *treap[K]) prev(key K) *node[K] { return t.kth(t.lowerBoundIndex(key) - 1) }

// 求 v 的后继（大于 key 的最小元素）
func (t *treap[K]) next(key K) *node[K] { return t.kth(t.upperBoundIndex(key)) }

// <= key 可以用 t.kth(t.upperBoundIndex(key)-1)
// >= key 可以用 t.kth(t.lowerBoundIndex(key))

func (t *treap[K]) find(key K) *node[K] {
	o := t.kth(t.lowerBoundIndex(key))
	if o == nil || o.key != key {
		return nil
	}
	return o
}

func newTreap[K cmp.Ordered]() *treap[K] {
	return &treap[K]{
		rd:         uint(time.Now().UnixNano())/2 + 1,
		comparator: cmp.Compare[K],
	}
}

func newTreapWith[K comparable](comp func(a, b K) int) *treap[K] {
	return &treap[K]{
		rd:         uint(time.Now().UnixNano())/2 + 1,
		comparator: comp,
	}
}

// 具体用法见本目录下的 test.go 文件
