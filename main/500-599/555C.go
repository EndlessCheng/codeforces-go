package main

import (
	"bufio"
	. "fmt"
	"io"
	"time"
)

// github.com/EndlessCheng/codeforces-go
type node55 struct {
	lr       [2]*node55
	priority uint
	key, end int
	up       bool
}

func (o *node55) cmp(b int) int {
	switch {
	case b < o.key:
		return 0
	case b > o.key:
		return 1
	default:
		return -1
	}
}

func (o *node55) rotate(d int) *node55 {
	x := o.lr[d^1]
	o.lr[d^1] = x.lr[d]
	x.lr[d] = o
	return x
}

type treap55 struct {
	rd   uint
	root *node55
}

func (t *treap55) fastRand() uint {
	t.rd ^= t.rd << 13
	t.rd ^= t.rd >> 17
	t.rd ^= t.rd << 5
	return t.rd
}

func (t *treap55) _put(o *node55, key, end int, up bool) *node55 {
	if o == nil {
		return &node55{priority: t.fastRand(), key: key, end: end, up: up}
	}
	if d := o.cmp(key); d >= 0 {
		o.lr[d] = t._put(o.lr[d], key, end, up)
		if o.lr[d].priority > o.priority {
			o = o.rotate(d ^ 1)
		}
	}
	return o
}

func (t *treap55) put(key, end int, up bool) { t.root = t._put(t.root, key, end, up) }

func (t *treap55) _delete(o *node55, key int) *node55 {
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

func (t *treap55) delete(key int) { t.root = t._delete(t.root, key) }

func (t *treap55) floor(key int) (floor *node55) {
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

func (t *treap55) lowerBound(key int) (lb *node55) {
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

func CF555C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	t := &treap55{rd: uint(time.Now().UnixNano())/2 + 1}
	var n, q, x, y, end int
	var dir []byte
	for Fscan(in, &n, &q); q > 0; q-- {
		Fscan(in, &x, &y, &dir)
		up := dir[0] == 'U'
		if up {
			if o := t.lowerBound(x); o == nil {
				end = 0
			} else if o.key == x {
				Fprintln(out, 0)
				continue
			} else if o.up {
				end = o.end
			} else {
				end = n + 1 - o.key
			}
			Fprintln(out, y-end)
		} else {
			if o := t.floor(x); o == nil {
				end = 0
			} else if o.key == x {
				Fprintln(out, 0)
				continue
			} else if o.up {
				end = o.key
			} else {
				end = o.end
			}
			Fprintln(out, x-end)
		}
		t.put(x, end, up)
	}
}

//func main() { CF555C(os.Stdin, os.Stdout) }
