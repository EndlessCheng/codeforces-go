package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

// https://www.luogu.com.cn/problem/P2887

// github.com/EndlessCheng/codeforces-go
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

func (t *treap) _put(o *node, key, val int) *node {
	if o == nil {
		return &node{priority: t.fastRand(), key: key, cnt: val}
	}
	if d := o.cmp(key); d >= 0 {
		o.lr[d] = t._put(o.lr[d], key, val)
		if o.lr[d].priority > o.priority {
			o = o.rotate(d ^ 1)
		}
	} else {
		o.cnt += val
	}
	return o
}

func (t *treap) put(key, val int) { t.root = t._put(t.root, key, val) }

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

func newTreap() *treap { return &treap{rd: 1} }

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
	return
}

func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, m, ans int
	Fscan(in, &n, &m)
	a := make([]struct{ l, r int }, n)
	for i := range a {
		Fscan(in, &a[i].l, &a[i].r)
	}
	sort.Slice(a, func(i, j int) bool { return a[i].r < a[j].r })
	t := newTreap()
	for ; m > 0; m-- {
		var v, c int
		Fscan(in, &v, &c)
		t.put(v, c)
	}
	for _, p := range a {
		if o := t.lowerBound(p.l); o != nil && o.key <= p.r {
			o.cnt--
			if o.cnt == 0 {
				t.delete(o.key)
			}
			ans++
		}
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
