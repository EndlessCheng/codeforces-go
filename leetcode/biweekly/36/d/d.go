package main

import "sort"

// github.com/EndlessCheng/codeforces-go
type node struct {
	lr       [2]*node
	priority uint
	key      int
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
		return &node{priority: t.fastRand(), key: key}
	}
	if d := o.cmp(key); d >= 0 {
		o.lr[d] = t._put(o.lr[d], key)
		if o.lr[d].priority > o.priority {
			o = o.rotate(d ^ 1)
		}
	}
	return o
}

func (t *treap) put(key int) { t.root = t._put(t.root, key) }

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

func (t *treap) lowerBound(key int) (lb *node) {
	for o := t.root; o != nil; {
		switch c := o.cmp(key); {
		case c == 0:
			lb = o
			o = o.lr[0]
		case c > 0:
			o = o.lr[1]
		default:
			return o
		}
	}
	if lb == nil {
		for o := t.root; o != nil; o = o.lr[0] {
			lb = o
		}
	}
	return
}

func newTreap() *treap { return &treap{rd: 1} }

func busiestServers(k int, starts, delta []int) (ans []int) {
	n := len(starts)
	type event struct{ p, d, i int }
	es := make([]event, 0, 2*n)
	for i, s := range starts {
		es = append(es, event{s, 1, i}, event{s + delta[i], -1, i})
	}
	sort.Slice(es, func(i, j int) bool { a, b := es[i], es[j]; return a.p < b.p || a.p == b.p && a.d < b.d })

	cnt := make([]int, k)
	maxCnt := 0
	idle := newTreap()
	for i := 0; i < k; i++ {
		idle.put(i)
	}
	server := make([]int, n)
	for _, e := range es {
		if e.d > 0 {
			o := idle.lowerBound(e.i % k)
			if o == nil {
				server[e.i] = -1
				continue
			}
			s := o.key
			if cnt[s]++; cnt[s] > maxCnt {
				maxCnt = cnt[s]
			}
			idle.delete(s)
			server[e.i] = s
		} else if s := server[e.i]; s >= 0 {
			idle.put(s)
		}
	}
	for i, v := range cnt {
		if v == maxCnt {
			ans = append(ans, i)
		}
	}
	return
}
