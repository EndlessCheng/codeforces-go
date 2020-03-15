package main

import "sort"

type node struct {
	lr       [2]*node
	priority uint
	key      int
	dupCnt   int
}

func (o *node) rotate(d int8) *node { x := o.lr[d^1]; o.lr[d^1] = x.lr[d]; x.lr[d] = o; return x }

type treap struct {
	rd   uint
	root *node
	ok   bool
}

func (t *treap) fastRand() uint {
	t.rd ^= t.rd << 13
	t.rd ^= t.rd >> 17
	t.rd ^= t.rd << 5
	return t.rd
}

func (t *treap) less(a, b int) int8 {
	switch {
	case a < b:
		return 0
	case a > b:
		return 1
	default:
		return -1
	}
}

func (t *treap) _put(o *node, key int) *node {
	if o == nil {
		return &node{priority: t.fastRand(), key: key, dupCnt: 1}
	}
	if d := t.less(key, o.key); d >= 0 {
		o.lr[d] = t._put(o.lr[d], key)
		if o.lr[d].priority > o.priority {
			o = o.rotate(d ^ 1)
		}
	} else {
		o.dupCnt++
	}
	return o
}

func (t *treap) put(key int) { t.root = t._put(t.root, key) }

func (t *treap) _delete(o *node, key int) *node {
	if o == nil {
		return nil
	}
	if d := t.less(key, o.key); d >= 0 {
		o.lr[d] = t._delete(o.lr[d], key)
	} else {
		t.ok = true
		if o.dupCnt > 1 {
			o.dupCnt--
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

func (t *treap) delete(key int) bool {
	t.ok = false
	t.root = t._delete(t.root, key)
	return t.ok
}

func (t *treap) max() (o *node) {
	if t.root == nil {
		return nil
	}
	for o = t.root; o.lr[1] != nil; o = o.lr[1] {
	}
	return
}

func newTreap() *treap { return &treap{rd: 1} }

func maxPerformance(n int, speed, efficiency []int, k int) (ans int) {
	type pair struct{ eff, spd int }
	ps := make([]pair, n)
	for i, s := range speed {
		ps[i] = pair{efficiency[i], s}
	}
	sort.Slice(ps, func(i, j int) bool { a, b := ps[i], ps[j]; return a.eff < b.eff || a.eff == b.eff && a.spd > b.spd })

	sort.Ints(speed)
	speedK := newTreap()
	speedSum := 0
	for i := n - 1; i >= n-k; i-- {
		speedK.put(speed[i])
		speedSum += speed[i]
	}
	speedLeft := newTreap()
	for i := n - k - 1; i >= 0; i-- {
		speedLeft.put(speed[i])
	}

	for _, p := range ps {
		spd := p.spd
		if !speedK.delete(spd) {
			speedLeft.delete(spd)
			continue
		}
		if val := p.eff * speedSum; val > ans {
			ans = val
		}
		speedSum -= spd
		if mx := speedLeft.max(); mx != nil {
			speedSum += mx.key
			speedK.put(mx.key)
			speedLeft.delete(mx.key)
		}
	}
	return ans % (1e9 + 7)
}
