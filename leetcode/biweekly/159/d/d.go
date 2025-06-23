package main

import (
	"cmp"
	"time"
)

// https://space.bilibili.com/206214

// 泛型 Treap 模板（set 版本，不含重复元素）
type nodeS[K comparable] struct {
	son      [2]*nodeS[K]
	priority uint
	key      K
	subSize  int
}

func (o *nodeS[K]) size() int {
	if o != nil {
		return o.subSize
	}
	return 0
}

func (o *nodeS[K]) maintain() {
	o.subSize = 1 + o.son[0].size() + o.son[1].size()
}

func (o *nodeS[K]) rotate(d int) *nodeS[K] {
	x := o.son[d^1]
	o.son[d^1] = x.son[d]
	x.son[d] = o
	o.maintain()
	x.maintain()
	return x
}

type treapS[K comparable] struct {
	rd         uint
	root       *nodeS[K]
	comparator func(a, b K) int
}

func (t *treapS[K]) fastRand() uint {
	t.rd ^= t.rd << 13
	t.rd ^= t.rd >> 17
	t.rd ^= t.rd << 5
	return t.rd
}

func (t *treapS[K]) size() int   { return t.root.size() }
func (t *treapS[K]) empty() bool { return t.size() == 0 }

func (t *treapS[K]) _put(o *nodeS[K], key K) *nodeS[K] {
	if o == nil {
		o = &nodeS[K]{priority: t.fastRand(), key: key}
	} else {
		c := t.comparator(key, o.key)
		if c != 0 {
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

func (t *treapS[K]) put(key K) { t.root = t._put(t.root, key) }

func (t *treapS[K]) _delete(o *nodeS[K], key K) *nodeS[K] {
	if o == nil {
		return nil
	}
	if c := t.comparator(key, o.key); c != 0 {
		d := 0
		if c > 0 {
			d = 1
		}
		o.son[d] = t._delete(o.son[d], key)
	} else {
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

func (t *treapS[K]) delete(key K) { t.root = t._delete(t.root, key) }

func (t *treapS[K]) min() *nodeS[K] { return t.kth(0) }
func (t *treapS[K]) max() *nodeS[K] { return t.kth(t.size() - 1) }

func (t *treapS[K]) lowerBoundIndex(key K) (kth int) {
	for o := t.root; o != nil; {
		c := t.comparator(key, o.key)
		if c < 0 {
			o = o.son[0]
		} else if c > 0 {
			kth += o.son[0].size() + 1
			o = o.son[1]
		} else {
			kth += o.son[0].size()
			break
		}
	}
	return
}

func (t *treapS[K]) upperBoundIndex(key K) (kth int) {
	for o := t.root; o != nil; {
		c := t.comparator(key, o.key)
		if c < 0 {
			o = o.son[0]
		} else if c > 0 {
			kth += o.son[0].size() + 1
			o = o.son[1]
		} else {
			kth += o.son[0].size() + 1
			break
		}
	}
	return
}

func (t *treapS[K]) kth(k int) (o *nodeS[K]) {
	if k < 0 || k >= t.root.size() {
		return
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
	return
}

func (t *treapS[K]) prev(key K) *nodeS[K] { return t.kth(t.lowerBoundIndex(key) - 1) }
func (t *treapS[K]) next(key K) *nodeS[K] { return t.kth(t.upperBoundIndex(key)) }

func (t *treapS[K]) find(key K) *nodeS[K] {
	o := t.kth(t.lowerBoundIndex(key))
	if o == nil || o.key != key {
		return nil
	}
	return o
}

func newSet[K cmp.Ordered]() *treapS[K] {
	return &treapS[K]{
		rd:         uint(time.Now().UnixNano()),
		comparator: cmp.Compare[K],
	}
}

func newSetWith[K comparable](comp func(a, b K) int) *treapS[K] {
	return &treapS[K]{
		rd:         uint(time.Now().UnixNano()),
		comparator: comp,
	}
}

func kthSmallest2(par []int, vals []int, queries [][]int) []int {
	n := len(par)
	g := make([][]int, n)
	for i := 1; i < n; i++ {
		p := par[i]
		g[p] = append(g[p], i)
	}

	type pair struct{ k, i int }
	qs := make([][]pair, n)
	for i, q := range queries {
		x, k := q[0], q[1]
		qs[x] = append(qs[x], pair{k, i})
	}

	ans := make([]int, len(queries))
	var dfs func(int, int) *treapS[int]
	dfs = func(x, xor int) *treapS[int] {
		xor ^= vals[x]
		set := newSet[int]()
		set.put(xor)
		for _, y := range g[x] {
			setY := dfs(y, xor)

			// 启发式合并：小集合并入大集合
			if setY.size() > set.size() {
				set, setY = setY, set
			}
			// 中序遍历 setY
			var f func(*nodeS[int])
			f = func(node *nodeS[int]) {
				if node == nil {
					return
				}
				f(node.son[0])
				set.put(node.key)
				f(node.son[1])
			}
			f(setY.root)
		}
		for _, p := range qs[x] {
			node := set.kth(p.k - 1)
			if node == nil {
				ans[p.i] = -1
			} else {
				ans[p.i] = node.key
			}
		}
		return set
	}
	dfs(0, 0)

	return ans
}
