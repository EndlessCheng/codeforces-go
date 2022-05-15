package main

import "time"

// github.com/EndlessCheng/codeforces-go
type CountIntervals struct {

}

type node struct {
	lr       [2]*node
	priority uint
	l,r        int
}

func (o *node) cmp(b int) int {
	switch {
	case b < o.l:
		return 0
	case b > o.l:
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

func (t *treap) _put(o *node, l,r int) *node {
	if o == nil {
		return &node{priority: t.fastRand(), l: l, r:r}
	}
	if d := o.cmp(l); d >= 0 {
		o.lr[d] = t._put(o.lr[d], l,r)
		if o.lr[d].priority > o.priority {
			o = o.rotate(d ^ 1)
		}
	}
	return o
}

func (t *treap) put(l,r int) { t.root = t._put(t.root, l,r) }

func (t *treap) _delete(o *node, key int) *node {
	if o == nil {
		return nil
	}
	if d := o.cmp(key); d >= 0 {
		o.lr[d] = t._delete(o.lr[d], key)
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
	return o
}

func (t *treap) delete(key int) { t.root = t._delete(t.root, key) }

func (t *treap) floor(key int) (prev *node) {
	for o := t.root; o != nil; {
		if o.cmp(key) <= 0 {
			o = o.lr[0]
		} else {
			prev = o
			o = o.lr[1]
		}
	}
	return // NOTE: check nil
}

func (t *treap) next(key int) (next *node) {
	for o := t.root; o != nil; {
		if o.cmp(key) == 0 {
			next = o
			o = o.lr[0]
		} else {
			o = o.lr[1]
		}
	}
	return // NOTE: check nil
}

func (t *treap) split(mid int) {
	if o := t.floor(mid); o != nil && o.l < mid && mid <= o.r {
		r := o.r
		o.r = mid - 1
		t.put(mid, r)
	}
}

func (t *treap) prepare(l, r int) {
	t.split(l)
	t.split(r + 1)
}

func newTreap() *treap { return &treap{rd: uint(time.Now().UnixNano())/2 + 1} }

var (
	t *treap
	sum int
)

func Constructor() CountIntervals {
	t = newTreap()
	sum = 0
	return CountIntervals{}
}

func (o *CountIntervals) Add(l, r int) {
	t.prepare(l, r)
	for o := t.next(l-1); o != nil && o.l <= r; o = t.next(o.l) {
		sum -= o.r-o.l+1
		t.delete(o.l)
	}
	sum += r-l+1
	t.put(l,r)
}

func (o *CountIntervals) Count() int { return sum }
