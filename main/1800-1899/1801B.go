package main

import (
	. "fmt"
	"io"
	"math/rand"
	"slices"
)

type node01 struct {
	lr       [2]*node01
	priority int
	key      int
}

func (o *node01) rotate(d int) *node01 {
	x := o.lr[d^1]
	o.lr[d^1] = x.lr[d]
	x.lr[d] = o
	return x
}

type treap01 struct {
	root *node01
}

func (t *treap01) _put(o *node01, key int) *node01 {
	if o == nil {
		return &node01{priority: rand.Int(), key: key}
	}
	if d := o.cmp(key); d >= 0 {
		o.lr[d] = t._put(o.lr[d], key)
		if o.lr[d].priority > o.priority {
			o = o.rotate(d ^ 1)
		}
	}
	return o
}

func (t *treap01) put(key int) { t.root = t._put(t.root, key) }

func (o *node01) cmp(key int) int {
	cur := o.key
	if key == cur {
		return -1
	}
	if key < cur {
		return 0
	}
	return 1
}

func (t *treap01) lowerBound(key int) (lb *node01) {
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

func (t *treap01) prev(key int) (prev *node01) {
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

func cf1801B(in io.Reader, out io.Writer) {
	const inf int = 1e18
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		type pair struct{ x, y int }
		a := make([]pair, n)
		for i := range a {
			Fscan(in, &a[i].x, &a[i].y)
		}
		slices.SortFunc(a, func(a, b pair) int { return a.x - b.x })
		suf := make([]int, n+1)
		suf[n] = -inf
		for i := n - 1; i > 0; i-- {
			suf[i] = max(suf[i+1], a[i].y)
		}

		ans := inf
		t := &treap01{}
		t.put(-inf)
		t.put(inf)
		for i, p := range a {
			x := p.x
			mx := suf[i+1]
			if mx >= x {
				ans = min(ans, mx-x)
			} else {
				ans = min(ans, x-max(t.prev(x).key, mx), t.lowerBound(x).key-x)
			}
			t.put(p.y)
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1801B(bufio.NewReader(os.Stdin), os.Stdout) }
