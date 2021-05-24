package main

import (
	"bufio"
	. "fmt"
	"io"
	"strings"
	"time"
)

// github.com/EndlessCheng/codeforces-go
type node58 struct {
	lr       [2]*node58
	priority uint
	l, r     int
	b        byte
}

func (o *node58) cmp(b int) int {
	switch {
	case b < o.l:
		return 0
	case b > o.l:
		return 1
	default:
		return -1
	}
}

func (o *node58) rotate(d int) *node58 {
	x := o.lr[d^1]
	o.lr[d^1] = x.lr[d]
	x.lr[d] = o
	return x
}

type treap58 struct {
	rd   uint
	root *node58
}

func (t *treap58) fastRand() uint {
	t.rd ^= t.rd << 13
	t.rd ^= t.rd >> 17
	t.rd ^= t.rd << 5
	return t.rd
}

func (t *treap58) _put(o *node58, l, r int, b byte) *node58 {
	if o == nil {
		return &node58{priority: t.fastRand(), l: l, r: r, b: b}
	}
	if d := o.cmp(l); d >= 0 {
		o.lr[d] = t._put(o.lr[d], l, r, b)
		if o.lr[d].priority > o.priority {
			o = o.rotate(d ^ 1)
		}
	} else {
		o.b = b
	}
	return o
}

func (t *treap58) put(l, r int, b byte) { t.root = t._put(t.root, l, r, b) }

func (t *treap58) _delete(o *node58, l int) *node58 {
	if o == nil {
		return nil
	}
	if d := o.cmp(l); d >= 0 {
		o.lr[d] = t._delete(o.lr[d], l)
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
		o.lr[d] = t._delete(o.lr[d], l)
	}
	return o
}

func (t *treap58) delete(l int) { t.root = t._delete(t.root, l) }

func (t *treap58) floor(key int) (floor *node58) {
	for o := t.root; o != nil; {
		switch c := o.cmp(key); {
		case c == 0:
			o = o.lr[0]
		case c > 0:
			floor = o
			o = o.lr[1]
		default:
			return o
		}
	}
	return
}

func (t *treap58) next(l int) (next *node58) {
	for o := t.root; o != nil; {
		if o.cmp(l) == 0 {
			next = o
			o = o.lr[0]
		} else {
			o = o.lr[1]
		}
	}
	return
}

func (t *treap58) split(mid int) {
	if o := t.floor(mid); o.l < mid && mid <= o.r {
		r, b := o.r, o.b
		o.r = mid - 1
		t.put(mid, r, b)
	}
}

func (t *treap58) prepare(l, r int) {
	t.split(l)
	t.split(r + 1)
}

func (t *treap58) sort(l, r int, inc bool) {
	t.prepare(l, r)
	cnt := [26]int{}
	for o := t.floor(l); o != nil && o.l <= r; o = t.next(o.l) {
		cnt[o.b] += o.r - o.l + 1
		t.delete(o.l)
	}
	if inc {
		for i, c := range cnt {
			if c > 0 {
				t.put(l, l+c-1, byte(i))
				l += c
			}
		}
	} else {
		for i, c := range cnt {
			if c > 0 {
				t.put(r-c+1, r, byte(i))
				r -= c
			}
		}
	}
}

func CF558E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, q, l, r int
	var inc bool
	var s []byte
	Fscan(in, &n, &q, &s)
	t := &treap58{rd: uint(time.Now().UnixNano())/2 + 1}
	for i, b := range s {
		t.put(i+1, i+1, b-'a')
	}
	for ; q > 0; q-- {
		Fscan(in, &l, &r, &inc)
		t.sort(l, r, inc)
	}
	var f func(*node58)
	f = func(o *node58) {
		if o == nil {
			return
		}
		f(o.lr[0])
		Fprint(out, strings.Repeat(string('a'+o.b), o.r-o.l+1))
		f(o.lr[1])
	}
	f(t.root)
}

//func main() { CF558E(os.Stdin, os.Stdout) }
