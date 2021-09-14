package main

import "sort"

// github.com/EndlessCheng/codeforces-go
type MKAverage struct{}

type node struct {
	lr       [2]*node
	priority uint
	key, cnt int
}

func (o *node) cmp(b int) int {
	switch {
	case b < o.key:
		return 0
	case b > o.key:
		return 1
	default:
		return -1
	}
}

func (o *node) rotate(d int) *node {
	x := o.lr[d^1]
	o.lr[d^1] = x.lr[d]
	x.lr[d] = o
	return x
}

type treap struct {
	rd   uint
	root *node
}

func (t *treap) fastRand() uint {
	t.rd ^= t.rd << 13
	t.rd ^= t.rd >> 17
	t.rd ^= t.rd << 5
	return t.rd
}

func (t *treap) _put(o *node, key int) *node {
	if o == nil {
		return &node{priority: t.fastRand(), key: key, cnt: 1}
	}
	if d := o.cmp(key); d >= 0 {
		o.lr[d] = t._put(o.lr[d], key)
		if o.lr[d].priority > o.priority {
			o = o.rotate(d ^ 1)
		}
	} else {
		o.cnt++
	}
	return o
}

func (t *treap) put(key int) {
	if t == mid {
		midSum += key
	}
	t.root = t._put(t.root, key)
}

func (t *treap) _delete(o *node, key int) *node {
	if o == nil {
		return nil
	}
	if d := o.cmp(key); d >= 0 {
		o.lr[d] = t._delete(o.lr[d], key)
	} else {
		if o.cnt > 1 {
			o.cnt--
		} else {
			if o.lr[1] == nil {
				return o.lr[0]
			}
			if o.lr[0] == nil {
				return o.lr[1]
			}
			d = 0
			if o.lr[0].priority > o.lr[1].priority {
				d = 1
			}
			o = o.rotate(d)
			o.lr[d] = t._delete(o.lr[d], key)
		}
	}
	return o
}

func (t *treap) delete(key int) {
	if t == mid {
		midSum -= key
	}
	t.root = t._delete(t.root, key)
}

func (t *treap) min() (min *node) {
	for o := t.root; o != nil; o = o.lr[0] {
		min = o
	}
	return
}

func (t *treap) max() (max *node) {
	for o := t.root; o != nil; o = o.lr[1] {
		max = o
	}
	return
}

func (t *treap) popMin() int {
	v := t.min().key
	t.delete(v)
	return v
}

func (t *treap) popMax() int {
	v := t.max().key
	t.delete(v)
	return v
}

func newTreap() *treap { return &treap{rd: 1} }

var (
	l, mid, r    *treap
	m, k, midSum int
	a            []int
)

func Constructor(M, K int) (_ MKAverage) {
	l, mid, r = newTreap(), newTreap(), newTreap()
	m, k, midSum = M, K, 0
	a = []int{}
	return
}

func (MKAverage) AddElement(num int) {
	a = append(a, num)
	if len(a) == m {
		b := append([]int(nil), a...)
		sort.Ints(b)
		for _, v := range b[:k] {
			l.put(v)
		}
		for _, v := range b[k : m-k] {
			mid.put(v)
		}
		for _, v := range b[m-k:] {
			r.put(v)
		}
	} else if len(a) > m {
		v := a[0]
		a = a[1:]
		if mx := l.max().key; v <= mx {
			l.delete(v)
			l.put(mid.popMin())
		} else if mi := r.min().key; v >= mi {
			r.delete(v)
			r.put(mid.popMax())
		} else {
			mid.delete(v)
		}

		v = num
		if mx := l.max().key; v <= mx {
			l.delete(mx)
			l.put(v)
			mid.put(mx)
		} else if mi := r.min().key; v >= mi {
			r.delete(mi)
			r.put(v)
			mid.put(mi)
		} else {
			mid.put(v)
		}
	}
}

func (MKAverage) CalculateMKAverage() int {
	if len(a) < m {
		return -1
	}
	return midSum / (m - 2*k)
}
