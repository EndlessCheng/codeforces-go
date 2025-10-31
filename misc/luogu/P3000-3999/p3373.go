package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// https://space.bilibili.com/206214
const mod373 = 571373

type pair373 struct{ k, c int }

type seg373 []struct {
	l, r int
	sum  int
	todo pair373
}

func (seg373) mergeInfo(a, b int) int {
	return (a + b) % mod373
}

var todoInit373 = pair373{1, 0}

func (seg373) mergeTodo(f, old pair373) pair373 {
	return pair373{f.k * old.k % mod373, (f.k*old.c + f.c) % mod373}
}

func (t seg373) apply(o int, f pair373) {
	cur := &t[o]

	sz := cur.r - cur.l + 1
	cur.sum = (f.k*cur.sum + f.c*sz) % mod373

	cur.todo = t.mergeTodo(f, cur.todo)
}

func (t seg373) maintain(o int) {
	t[o].sum = t.mergeInfo(t[o<<1].sum, t[o<<1|1].sum)
}

func (t seg373) spread(o int) {
	f := t[o].todo
	if f == todoInit373 {
		return
	}
	t.apply(o<<1, f)
	t.apply(o<<1|1, f)
	t[o].todo = todoInit373
}

func (t seg373) build(a []int, o, l, r int) {
	t[o].l, t[o].r = l, r
	t[o].todo = todoInit373
	if l == r {
		t[o].sum = a[l]
		return
	}
	m := (l + r) >> 1
	t.build(a, o<<1, l, m)
	t.build(a, o<<1|1, m+1, r)
	t.maintain(o)
}

func (t seg373) update(o, l, r int, f pair373) {
	if l <= t[o].l && t[o].r <= r {
		t.apply(o, f)
		return
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if l <= m {
		t.update(o<<1, l, r, f)
	}
	if m < r {
		t.update(o<<1|1, l, r, f)
	}
	t.maintain(o)
}

func (t seg373) query(o, l, r int) int {
	if l <= t[o].l && t[o].r <= r {
		return t[o].sum
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.query(o<<1, l, r)
	}
	if l > m {
		return t.query(o<<1|1, l, r)
	}
	return t.mergeInfo(t.query(o<<1, l, r), t.query(o<<1|1, l, r))
}

func p3373(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, q, op, l, r, k int
	Fscan(in, &n, &q, &l)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	t := make(seg373, 2<<bits.Len(uint(n-1)))
	t.build(a, 1, 0, n-1)
	for ; q > 0; q-- {
		Fscan(in, &op, &l, &r)
		l--
		r--
		if op == 1 {
			Fscan(in, &k)
			t.update(1, l, r, pair373{k, 0})
		} else if op == 2 {
			Fscan(in, &k)
			t.update(1, l, r, pair373{1, k})
		} else {
			Fprintln(out, t.query(1, l, r))
		}
	}
}

//func main() { p3373(bufio.NewReader(os.Stdin), os.Stdout) }
