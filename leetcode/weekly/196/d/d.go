package main

type node struct {
	lr       [2]*node
	priority uint
	key      int
	sz       int
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

func (o *node) size() int {
	if o != nil {
		return o.sz
	}
	return 0
}

func (o *node) pushUp() { o.sz = 1 + o.lr[0].size() + o.lr[1].size() }

func (o *node) rotate(d int8) *node {
	x := o.lr[d^1]
	o.lr[d^1] = x.lr[d]
	x.lr[d] = o
	o.pushUp()
	x.pushUp()
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
		return &node{priority: t.fastRand(), key: key, sz: 1}
	}
	if d := o.cmp(key); d >= 0 {
		o.lr[d] = t._put(o.lr[d], key)
		if o.lr[d].priority > o.priority {
			o = o.rotate(d ^ 1)
		}
	}
	o.pushUp()
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
	o.pushUp()
	return o
}

func (t *treap) delete(key int) { t.root = t._delete(t.root, key) }

func newTreap() *treap { return &treap{rd: 1} }

func (t *treap) cntLess(key int) (cnt int) {
	for o := t.root; o != nil; {
		switch c := o.cmp(key); {
		case c == 0:
			o = o.lr[0]
		case c > 0:
			cnt += 1 + o.lr[0].size()
			o = o.lr[1]
		default:
			cnt += o.lr[0].size()
			return
		}
	}
	return
}

func minInteger(s string, k int) (ans_ string) {
	n := len(s)
	pos := [10][]int{}
	for i, c := range s {
		c -= '0'
		pos[c] = append(pos[c], i)
	}

	ans := make([]byte, 0, n)
	t := newTreap()
o:
	for k > 0 {
		for i := 0; i < 9; i++ {
			if len(pos[i]) == 0 {
				continue
			}
			p := pos[i][0]
			rightShift := t.root.size() - t.cntLess(p)
			cost := p + rightShift - len(ans)
			if cost > k {
				continue
			}
			if cost > 0 {
				k -= cost
				t.put(p)
			}
			ans = append(ans, '0'+byte(i))
			pos[i] = pos[i][1:]
			continue o
		}
		break
	}

	left := make([]byte, n)
	for i, ps := range pos {
		for _, p := range ps {
			left[p] = '0' + byte(i)
		}
	}
	for _, b := range left {
		if b > 0 {
			ans = append(ans, b)
		}
	}

	if k > 0 {
		for i := k; i < n; i++ {
			if ans[i-k] > ans[i] {
				ans[i], ans[i-k] = ans[i-k], ans[i]
				break
			}
		}
	}

	return string(ans)
}
