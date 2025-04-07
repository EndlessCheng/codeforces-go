package main

import (
	"cmp"
	"time"
)

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

func minimumPairRemoval2(a []int) (ans int) {
	type pair struct{ s, i int }
	ps := newSetWith[pair](func(a, b pair) int { return cmp.Or(a.s-b.s, a.i-b.i) })
	dec := 0
	add := func(x, y, i int) {
		if x > y {
			dec++
		}
		ps.put(pair{x + y, i})
	}
	del := func(x, y, i int) {
		if x > y {
			dec--
		}
		ps.delete(pair{x + y, i})
	}

	n := len(a)
	for i := 1; i < n; i++ {
		add(a[i-1], a[i], i)
	}
	idx := newSet[int]()
	for i := range n {
		idx.put(i)
	}

	for dec > 0 {
		ans++

		p := ps.min().key
		s, i := p.s, p.i
		k := idx.lowerBoundIndex(i)

		l := idx.kth(k - 1).key
		del(a[l], a[i], i)

		if k > 1 {
			ll := idx.kth(k - 2).key
			del(a[ll], a[l], l)
			add(a[ll], s, l)
		}

		if k+1 < idx.size() {
			r := idx.kth(k + 1).key
			del(a[i], a[r], r)
			add(s, a[r], r)
		}

		a[l] = s
		idx.delete(i)
	}
	return
}
