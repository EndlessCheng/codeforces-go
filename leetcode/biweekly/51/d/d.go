package main

import "sort"

// github.com/EndlessCheng/codeforces-go
type node struct {
	ch       [2]*node
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
	x := o.ch[d^1]
	o.ch[d^1] = x.ch[d]
	x.ch[d] = o
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
		o.ch[d] = t._put(o.ch[d], key)
		if o.ch[d].priority > o.priority {
			o = o.rotate(d ^ 1)
		}
	}
	return o
}

func (t *treap) put(key int) { t.root = t._put(t.root, key) }

func (t *treap) prev(key int) (prev *node) {
	for o := t.root; o != nil; {
		if o.cmp(key) <= 0 {
			o = o.ch[0]
		} else {
			prev = o
			o = o.ch[1]
		}
	}
	return
}

func (t *treap) lowerBound(key int) (lb *node) {
	for o := t.root; o != nil; {
		switch c := o.cmp(key); {
		case c == 0:
			lb = o
			o = o.ch[0]
		case c > 0:
			o = o.ch[1]
		default:
			return o
		}
	}
	return
}

func closestRoom(a, qs [][]int) []int {
	ans := make([]int, len(qs))
	sort.Slice(a, func(i, j int) bool { return a[i][1] > a[j][1] })
	for i := range qs {
		qs[i] = append(qs[i], i)
	}
	sort.Slice(qs, func(i, j int) bool { return qs[i][1] > qs[j][1] })
	t := &treap{rd: 1}
	i, n := 0, len(a)
	for _, q := range qs {
		for ; i < n && a[i][1] >= q[1]; i++ {
			t.put(a[i][0])
		}
		tar := q[0]
		res := -1
		if o := t.prev(tar); o != nil {
			res = o.key
		}
		if o := t.lowerBound(tar); o != nil && (res < 0 || o.key-tar < tar-res) {
			res = o.key
		}
		ans[q[2]] = res
	}
	return ans
}
