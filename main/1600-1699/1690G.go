package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
type node90 struct {
	lr       [2]*node90
	priority uint
	key      int
}

func (o *node90) cmp(b int) int {
	switch {
	case b < o.key:
		return 0
	case b > o.key:
		return 1
	default:
		return -1
	}
}

func (o *node90) rotate(d int) *node90 {
	x := o.lr[d^1]
	o.lr[d^1] = x.lr[d]
	x.lr[d] = o
	return x
}

type treap90 struct {
	rd   uint
	root *node90
}

func (t *treap90) fastRand() uint {
	t.rd ^= t.rd << 13
	t.rd ^= t.rd >> 17
	t.rd ^= t.rd << 5
	return t.rd
}

func (t *treap90) _put(o *node90, key int) *node90 {
	if o == nil {
		return &node90{priority: t.fastRand(), key: key}
	}
	if d := o.cmp(key); d >= 0 {
		o.lr[d] = t._put(o.lr[d], key)
		if o.lr[d].priority > o.priority {
			o = o.rotate(d ^ 1)
		}
	}
	return o
}

func (t *treap90) put(key int) { t.root = t._put(t.root, key) }

func (t *treap90) _delete(o *node90, key int) *node90 {
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

func (t *treap90) delete(key int) { t.root = t._delete(t.root, key) }

func (t *treap90) prev(key int) (prev *node90) {
	for o := t.root; o != nil; {
		if o.cmp(key) <= 0 {
			o = o.lr[0]
		} else {
			prev = o
			o = o.lr[1]
		}
	}
	return
}

func (t *treap90) next(key int) (next *node90) {
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

func CF1690G(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, m, i, d int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		ans := 0
		t := &treap90{rd: 1}
		mi := a[0] + 1
		for i, v := range a {
			if v < mi {
				ans++
				t.put(i)
				mi = v
			}
		}
		for ; m > 0; m-- {
			Fscan(in, &i, &d)
			i--
			v := a[i] - d
			if o := t.prev(i); o != nil && a[i] >= a[o.key] && v < a[o.key] {
				ans++
				t.put(i)
			}
			for o := t.next(i); o != nil && a[o.key] >= v; o = t.next(i) {
				ans--
				t.delete(o.key)
			}
			a[i] = v
			Fprint(out, ans, " ")
		}
		Fprintln(out)
	}
}

//func main() { CF1690G(os.Stdin, os.Stdout) }
