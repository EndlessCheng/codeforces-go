package main

import (
	"bufio"
	. "fmt"
	"io"
	"time"
)

// github.com/EndlessCheng/codeforces-go
type node28 struct {
	lr       [2]*node28
	priority uint
	key      int
}

func (o *node28) cmp(b int) int {
	switch {
	case b < o.key:
		return 0
	case b > o.key:
		return 1
	default:
		return -1
	}
}

func (o *node28) rotate(d int) *node28 {
	x := o.lr[d^1]
	o.lr[d^1] = x.lr[d]
	x.lr[d] = o
	return x
}

type treap28 struct {
	rd   uint
	root *node28
}

func (t *treap28) fastRand() uint {
	t.rd ^= t.rd << 13
	t.rd ^= t.rd >> 17
	t.rd ^= t.rd << 5
	return t.rd
}

func (t *treap28) _put(o *node28, key int) *node28 {
	if o == nil {
		return &node28{priority: t.fastRand(), key: key}
	}
	if d := o.cmp(key); d >= 0 {
		o.lr[d] = t._put(o.lr[d], key)
		if o.lr[d].priority > o.priority {
			o = o.rotate(d ^ 1)
		}
	}
	return o
}

func (t *treap28) put(key int) { t.root = t._put(t.root, key) }

func (t *treap28) _delete(o *node28, key int) *node28 {
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

func (t *treap28) delete(key int) { t.root = t._delete(t.root, key) }

func (t *treap28) lowerBound(key int) (lb *node28) {
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

func (t *treap28) prev(key int) (prev *node28) {
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

func CF1528C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	rd := uint(time.Now().UnixNano())/2 + 1

	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		g1 := make([][]int, n)
		g2 := make([][]int, n)
		for w := 1; w < n; w++ {
			Fscan(in, &v)
			v--
			g1[v] = append(g1[v], w)
		}
		for w := 1; w < n; w++ {
			Fscan(in, &v)
			v--
			g2[v] = append(g2[v], w)
		}

		l := make([]int, n)
		r := make([]int, n)
		at := make([]int, n+1)
		ts := 0
		var f2 func(int)
		f2 = func(v int) {
			ts++
			l[v] = ts
			at[ts] = v
			for _, w := range g2[v] {
				f2(w)
			}
			r[v] = ts
		}
		f2(0)

		ans, sz := 0, 0
		t := &treap28{rd: rd}
		var f func(int)
		f = func(v int) {
			lb := t.lowerBound(l[v])
			if lb == nil || lb.key > r[v] { // 子树没有节点在 t 中
				t.put(l[v])
				sz++
				defer func() { t.delete(l[v]); sz-- }()
				if o := t.prev(l[v]); o != nil {
					if pa := at[o.key]; l[pa] < l[v] && l[v] <= r[pa] { // 有父节点在 t 中
						t.delete(l[pa])
						sz--
						defer func() { t.put(l[pa]); sz++ }()
					}
				}
				if sz > ans {
					ans = sz
				}
			}
			for _, w := range g1[v] {
				f(w)
			}
		}
		f(0)
		Fprintln(out, ans)
	}
}

//func main() { CF1528C(os.Stdin, os.Stdout) }
