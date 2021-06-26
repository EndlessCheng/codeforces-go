package main

import (
	"time"
)

// github.com/EndlessCheng/codeforces-go
type MovieRentingSystem struct{}
type pair struct{ shop, movie int }
type entry struct{ price, shop, movie int }

const mx int = 1e4

var (
	movPrice map[pair]int
	has      [mx + 1]*treap
	rent     *treap
)

func Constructor(_ int, entries [][]int) (_ MovieRentingSystem) {
	movPrice = map[pair]int{}
	has = [mx + 1]*treap{}
	for i := 1; i <= mx; i++ {
		has[i] = newTreap()
	}
	rent = newTreap()
	for _, e := range entries {
		shop, movie, price := e[0], e[1], e[2]
		movPrice[pair{shop, movie}] = price
		has[movie].put(entry{price, shop, movie})
	}
	return
}

func (MovieRentingSystem) Search(movie int) (ans []int) {
	has[movie].foreach(func(o *node) bool {
		ans = append(ans, o.key.shop)
		return len(ans) == 5
	})
	return
}

func (MovieRentingSystem) Rent(shop, movie int) {
	e := entry{movPrice[pair{shop, movie}], shop, movie}
	has[movie].delete(e)
	rent.put(e)
}

func (MovieRentingSystem) Drop(shop, movie int) {
	e := entry{movPrice[pair{shop, movie}], shop, movie}
	rent.delete(e)
	has[movie].put(e)
}

func (MovieRentingSystem) Report() (ans [][]int) {
	rent.foreach(func(o *node) bool {
		ans = append(ans, []int{o.key.shop, o.key.movie})
		return len(ans) == 5
	})
	return
}

// 以下为 treap 模板

type node struct {
	ch       [2]*node
	priority uint
	key      entry
}

func less(a, b entry) bool {
	if a.price != b.price {
		return a.price < b.price
	}
	if a.shop != b.shop {
		return a.shop < b.shop
	}
	return a.movie < b.movie
}

func (o *node) cmp(b entry) int {
	switch {
	case b == o.key:
		return -1
	case less(b, o.key):
		return 0
	default:
		return 1
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

func (t *treap) _put(o *node, key entry) *node {
	if o == nil {
		return &node{priority: t.fastRand(), key: key}
	}
	d := o.cmp(key)
	o.ch[d] = t._put(o.ch[d], key)
	if o.ch[d].priority > o.priority {
		o = o.rotate(d ^ 1)
	}
	return o
}

func (t *treap) put(key entry) { t.root = t._put(t.root, key) }

func (t *treap) _delete(o *node, key entry) *node {
	if d := o.cmp(key); d >= 0 {
		o.ch[d] = t._delete(o.ch[d], key)
	} else {
		if o.ch[1] == nil {
			return o.ch[0]
		}
		if o.ch[0] == nil {
			return o.ch[1]
		}
		d = 0
		if o.ch[0].priority > o.ch[1].priority {
			d = 1
		}
		o = o.rotate(d)
		o.ch[d] = t._delete(o.ch[d], key)
	}
	return o
}

func (t *treap) delete(key entry) { t.root = t._delete(t.root, key) }

func (t *treap) foreach(do func(o *node) (Break bool)) {
	var f func(*node) bool
	f = func(o *node) bool {
		return o != nil && (f(o.ch[0]) || do(o) || f(o.ch[1]))
	}
	f(t.root)
}

func newTreap() *treap { return &treap{rd: uint(time.Now().UnixNano())/2 + 1} }
