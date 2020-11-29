package main

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

func (t *treap) max() int {
	var max *node
	for o := t.root; o != nil; o = o.lr[1] {
		max = o
	}
	return max.key
}

func (t *treap) min() int {
	var min *node
	for o := t.root; o != nil; o = o.lr[0] {
		min = o
	}
	return min.key
}

func minimumDeviation(a []int) (ans int) {
	t := &treap{rd: 1}
	for _, v := range a {
		t.put(v << (v & 1))
	}
	ans = 1e9
	for {
		mx := t.max()
		ans = min(ans, mx-t.min())
		if mx&1 > 0 {
			return
		}
		t.delete(mx)
		t.put(mx >> 1)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
