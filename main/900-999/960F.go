package main

import (
	"bufio"
	. "fmt"
	"io"
	"time"
)

// github.com/EndlessCheng/codeforces-go
type node60 struct {
	lr       [2]*node60
	priority uint
	key, val int
}

func (o *node60) cmp(b int) int {
	switch {
	case b < o.key:
		return 0
	case b > o.key:
		return 1
	default:
		return -1
	}
}

func (o *node60) rotate(d int) *node60 {
	x := o.lr[d^1]
	o.lr[d^1] = x.lr[d]
	x.lr[d] = o
	return x
}

type treap60 struct {
	rd   uint
	root *node60
}

func (t *treap60) fastRand() uint {
	t.rd ^= t.rd << 13
	t.rd ^= t.rd >> 17
	t.rd ^= t.rd << 5
	return t.rd
}

func (t *treap60) _put(o *node60, key, val int) *node60 {
	if o == nil {
		return &node60{priority: t.fastRand(), key: key, val: val}
	}
	if d := o.cmp(key); d >= 0 {
		o.lr[d] = t._put(o.lr[d], key, val)
		if o.lr[d].priority > o.priority {
			o = o.rotate(d ^ 1)
		}
	}
	return o
}

func (t *treap60) put(key, val int) { t.root = t._put(t.root, key, val) }

func (t *treap60) _delete(o *node60, key int) *node60 {
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

func (t *treap60) delete(key int) { t.root = t._delete(t.root, key) }

func (t *treap60) lowerBound(key int) (lb *node60) {
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

func (t *treap60) prev(key int) (prev *node60) {
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

func CF960F(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, m, v, w, wt, ans int
	Fscan(in, &n, &m)
	ts := make([]*treap60, n)
	rd := uint(time.Now().UnixNano())/2 + 1
	for i := range ts {
		ts[i] = &treap60{rd: rd}
	}
	for ; m > 0; m-- {
		Fscan(in, &v, &w, &wt)
		v--
		w--
		res := 1
		if o := ts[v].prev(wt + 1); o != nil {
			res = o.val + 1
		}
		if res > ans {
			ans = res
		}
		for {
			o := ts[w].lowerBound(wt)
			if o == nil || o.val > res {
				break
			}
			ts[w].delete(o.key)
		}
		if o := ts[w].lowerBound(wt); o != nil && o.key == wt {
			continue
		}
		if o := ts[w].prev(wt); o != nil && o.val >= res {
			continue
		}
		ts[w].put(wt, res)
	}
	Fprint(out, ans)
}

//func main() { CF960F(os.Stdin, os.Stdout) }
