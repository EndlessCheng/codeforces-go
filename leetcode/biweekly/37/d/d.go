package main

// github.com/EndlessCheng/codeforces-go
const mod int = 1e9 + 7

type node struct{ l, r, mulTodo, addTodo, s int }
type seg []node

func (t seg) maintain(o int) {
	t[o].s = (t[o<<1].s + t[o<<1|1].s) % mod
}
func (t seg) build(o, l, r int) {
	t[o] = node{l, r, 1, 0, 0}
	if l == r {
		return
	}
	m := (l + r) >> 1
	t.build(o<<1, l, m)
	t.build(o<<1|1, m+1, r)
	t.maintain(o)
}
func (t seg) doMul(o, v int) {
	to := &t[o]
	to.mulTodo = to.mulTodo * v % mod
	to.addTodo = to.addTodo * v % mod
	to.s = to.s * v % mod
}
func (t seg) doAdd(o, v int) {
	to := &t[o]
	to.s = (to.s + (to.r-to.l+1)*v) % mod
	to.addTodo = (to.addTodo + v) % mod
}
func (t seg) spread(o int) {
	if v := t[o].mulTodo; v != 1 {
		t.doMul(o<<1, v)
		t.doMul(o<<1|1, v)
		t[o].mulTodo = 1
	}
	if v := t[o].addTodo; v != 0 {
		t.doAdd(o<<1, v)
		t.doAdd(o<<1|1, v)
		t[o].addTodo = 0
	}
}
func (t seg) update(o, l, r, v int, mul bool) {
	if l <= t[o].l && t[o].r <= r {
		if mul {
			t.doMul(o, v)
		} else {
			t.doAdd(o, v)
		}
		return
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if l <= m {
		t.update(o<<1, l, r, v, mul)
	}
	if m < r {
		t.update(o<<1|1, l, r, v, mul)
	}
	t.maintain(o)
}
func (t seg) query(o, l, r int) int {
	if l <= t[o].l && t[o].r <= r {
		return t[o].s
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.query(o<<1, l, r)
	}
	if l > m {
		return t.query(o<<1|1, l, r)
	}
	return (t.query(o<<1, l, r) + t.query(o<<1|1, l, r)) % mod
}

type Fancy struct{}

const n int = 1e5

func Constructor() (f Fancy) {
	t.build(1, 1, n)
	cnt = 0
	return
}

var t = make(seg, 4*n)
var cnt int

func (Fancy) Append(v int) {
	cnt++
	t.update(1, cnt, cnt, v, false)
}

func (Fancy) AddAll(v int) {
	t.update(1, 1, cnt, v, false)
}

func (Fancy) MultAll(v int) {
	t.update(1, 1, cnt, v, true)
}

func (Fancy) GetIndex(idx int) int {
	idx++
	if idx > cnt {
		return -1
	}
	return t.query(1, idx, idx)
}
