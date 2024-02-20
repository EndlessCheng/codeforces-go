package main

import (
	"bufio"
	. "fmt"
	"io"
	"time"
)

// https://space.bilibili.com/206214
type node98 struct {
	lr       [2]*node98
	priority uint
	key      int
	value    map[int]bool
}

func (o *node98) rotate(d int) *node98 {
	x := o.lr[d^1]
	o.lr[d^1] = x.lr[d]
	x.lr[d] = o
	return x
}

type treap98 struct {
	rd   uint
	root *node98
}

func (t *treap98) fastRand() uint {
	t.rd ^= t.rd << 13
	t.rd ^= t.rd >> 17
	t.rd ^= t.rd << 5
	return t.rd
}

func (t *treap98) _put(o *node98, key int, value int) *node98 {
	if o == nil {
		return &node98{priority: t.fastRand(), key: key, value: map[int]bool{value: true}}
	}
	if d := o.cmp(key); d >= 0 {
		o.lr[d] = t._put(o.lr[d], key, value)
		if o.lr[d].priority > o.priority {
			o = o.rotate(d ^ 1)
		}
	} else {
		o.value[value] = true
	}
	return o
}

func (t *treap98) put(key int, value int) { t.root = t._put(t.root, key, value) }

func (t *treap98) _delete(o *node98, key int) *node98 {
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

func (t *treap98) delete(key int) { t.root = t._delete(t.root, key) }

func (o *node98) cmp(a int) int {
	b := o.key
	if a == b {
		return -1
	}
	if a < b {
		return 0
	}
	return 1
}

func (t *treap98) min() (min *node98) {
	for o := t.root; o != nil; o = o.lr[0] {
		min = o
	}
	return
}

func (t *treap98) get(key int) *node98 {
	for o := t.root; o != nil; {
		if c := o.cmp(key); c >= 0 {
			o = o.lr[c]
		} else {
			return o
		}
	}
	return nil
}

func cf1198B_treap(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, q, op int
	Fscan(in, &n)
	t := &treap98{rd: uint(time.Now().UnixNano())/2 + 1}
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		t.put(a[i], i)
	}
	Fscan(in, &q)
	qs := make([]struct{ p, x int }, q)
	mx := 0
	for i := range qs {
		Fscan(in, &op)
		if op == 1 {
			Fscan(in, &qs[i].p, &qs[i].x)
		} else {
			Fscan(in, &qs[i].x)
			mx = max(mx, qs[i].x)
		}
	}

	done := false
	for _, q := range qs {
		x := q.x
		if q.p > 0 {
			i := q.p - 1
			o := t.get(a[i])
			if o != nil {
				delete(o.value, i)
			}
			a[i] = x
			t.put(x, i)
		} else if done || x == mx {
			done = true
			for {
				o := t.min()
				if o == nil || o.key >= x {
					break
				}
				for i := range o.value {
					a[i] = x
					t.put(x, i)
				}
				t.delete(o.key)
			}
		}
	}
	for _, v := range a {
		Fprint(out, v, " ")
	}
}

//func main() { cf1198B(os.Stdin, os.Stdout) }
