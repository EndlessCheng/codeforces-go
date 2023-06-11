package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
	"time"
)

// https://space.bilibili.com/206214
type node struct {
	lr       [2]*node
	priority uint
	key      int
	keyCnt   int
	subCnt   int
}

func (o *node) size() int {
	if o != nil {
		return o.subCnt // 汇总
	}
	return 0
}

func (o *node) maintain() {
	o.subCnt = o.keyCnt + o.lr[0].size() + o.lr[1].size()
}

func (o *node) rotate(d int) *node {
	x := o.lr[d^1]
	o.lr[d^1] = x.lr[d]
	x.lr[d] = o
	o.maintain()
	x.maintain()
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

func (t *treap) size() int { return t.root.size() }

func (t *treap) _put(o *node, key int) *node {
	if o == nil {
		o = &node{priority: t.fastRand(), key: key, keyCnt: 1}
	} else if d := o.cmp(key); d >= 0 {
		o.lr[d] = t._put(o.lr[d], key)
		if o.lr[d].priority > o.priority {
			o = o.rotate(d ^ 1)
		}
	} else {
		o.keyCnt++
	}
	o.maintain()
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
		if o.keyCnt > 1 {
			o.keyCnt--
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
	o.maintain()
	return o
}

func (t *treap) delete(key int) { t.root = t._delete(t.root, key) }

func newTreap() *treap { return &treap{rd: uint(time.Now().UnixNano())/2 + 1} }

func (o *node) cmp(a int) int {
	b := o.key
	if a == b {
		return -1
	}
	if a < b {
		return 0
	}
	return 1
}

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
	return // NOTE: check nil
}

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m int
	Fscan(in, &n, &m)
	a := make([]struct{ x, y int }, n)
	for i := range a {
		Fscan(in, &a[i].x)
	}
	for i := range a {
		Fscan(in, &a[i].y)
	}
	sort.Slice(a, func(i, j int) bool { return a[i].x > a[j].x })
	b := make([]struct{ x, y int }, m)
	for i := range b {
		Fscan(in, &b[i].x)
	}
	for i := range b {
		Fscan(in, &b[i].y)
	}
	sort.Slice(b, func(i, j int) bool { return b[i].x > b[j].x })

	t := newTreap()
	i := 0
	for _, p := range a {
		for ; i < m && b[i].x >= p.x; i++ {
			t.put(b[i].y)
		}
		o := t.lowerBound(p.y)
		if o == nil {
			Fprint(out, "No")
			return
		}
		t.delete(o.key)
	}
	Fprint(out, "Yes")
}

func main() { run(os.Stdin, os.Stdout) }
