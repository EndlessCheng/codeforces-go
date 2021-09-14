package main

type node struct {
	lr       [2]*node
	priority uint
	key      int
}

func (o *node) cmp(b int) int8 {
	switch {
	case b < o.key:
		return 0
	case b > o.key:
		return 1
	default:
		return -1
	}
}

func (o *node) rotate(d int8) *node {
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

func (t *treap) next(key int) (next *node) {
	for o := t.root; o != nil; {
		if o.cmp(key) == 0 {
			next = o
			o = o.lr[0]
		} else {
			o = o.lr[1]
		}
	}
	return
}

func avoidFlood(a []int) (ans []int) {
	ans = make([]int, len(a))
	lpos := map[int]int{}
	zpos := &treap{rd: 1}
	for i, lake := range a {
		if lake == 0 {
			ans[i] = 1
			zpos.put(i)
			continue
		}
		ans[i] = -1
		if l, ok := lpos[lake]; ok {
			p := zpos.next(l)
			if p == nil {
				return []int{}
			}
			ans[p.key] = lake
			zpos.delete(p.key)
		}
		lpos[lake] = i
	}
	return
}
