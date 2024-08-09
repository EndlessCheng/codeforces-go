package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
	"os"
)

// https://space.bilibili.com/206214
const mod = 998244353
var todoInit pair

type data struct{ sa, sb, sab int }
type pair struct{ a, b int }
type seg []struct {
	l, r int
	d    data
	todo pair
}

func mergeInfo(a, b data) (c data) {
	c.sa = (a.sa + b.sa) % mod
	c.sb = (a.sb + b.sb) % mod
	c.sab = (a.sab + b.sab) % mod
	return
}

func (t seg) do(O int, p pair) {
	o := &t[O]

	sz := o.r - o.l + 1
	if p.a > 0 {
		o.d.sa = (o.d.sa + sz*p.a) % mod
		o.d.sab = (o.d.sab + o.d.sb*p.a) % mod
	}

	if p.b > 0 {
		o.d.sb = (o.d.sb + sz*p.b) % mod
		o.d.sab = (o.d.sab + o.d.sa*p.b) % mod
	}

	o.todo.a = (o.todo.a + p.a) % mod
	o.todo.b = (o.todo.b + p.b) % mod
}

func (t seg) spread(o int) {
	if v := t[o].todo; v != todoInit {
		t.do(o<<1, v)
		t.do(o<<1|1, v)
		t[o].todo = todoInit
	}
}

func (t seg) build(a [][2]int, o, l, r int) {
	t[o].l, t[o].r = l, r
	t[o].todo = todoInit
	if l == r {
		t[o].d.sa = a[l-1][0]
		t[o].d.sb = a[l-1][1]
		t[o].d.sab = a[l-1][0] * a[l-1][1] % mod
		return
	}
	m := (l + r) >> 1
	t.build(a, o<<1, l, m)
	t.build(a, o<<1|1, m+1, r)
	t.maintain(o)
}

func (t seg) update(o, l, r int, v pair) {
	if l <= t[o].l && t[o].r <= r {
		t.do(o, v)
		return
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if l <= m {
		t.update(o<<1, l, r, v)
	}
	if m < r {
		t.update(o<<1|1, l, r, v)
	}
	t.maintain(o)
}

func (t seg) maintain(o int) {
	t[o].d = mergeInfo(t[o<<1].d, t[o<<1|1].d)
}

func (t seg) query(o, l, r int) int {
	if l <= t[o].l && t[o].r <= r {
		return t[o].d.sab
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.query(o<<1, l, r)
	}
	if l > m {
		return t.query(o<<1|1, l, r)
	}
	return t.query(o<<1, l, r) + t.query(o<<1|1, l, r)
}

func run(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, q, op, l, r, x int
	Fscan(in, &n, &q)
	a := make([][2]int, n)
	for i := range a {
		Fscan(in, &a[i][0])
	}
	for i := range a {
		Fscan(in, &a[i][1])
	}
	t := make(seg, 2<<bits.Len(uint(n-1)))
	t.build(a, 1, 1, n)
	for ; q > 0; q-- {
		Fscan(in, &op, &l, &r)
		if op == 3 {
			Fprintln(out, t.query(1, l, r)%mod)
		} else {
			Fscan(in, &x)
			if op == 1 {
				t.update(1, l, r, pair{x, 0})
			} else {
				t.update(1, l, r, pair{0, x})
			}
		}
	}
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
