package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// https://space.bilibili.com/206214
type seg95 []struct{ l, r, min, todo int }

func (t seg95) do(o, v int) {
	t[o].min += v
	t[o].todo += v
}

func (t seg95) spread(o int) {
	if v := t[o].todo; v != 0 {
		t.do(o<<1, v)
		t.do(o<<1|1, v)
		t[o].todo = 0
	}
}

func (t seg95) build(o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		return
	}
	m := (l + r) >> 1
	t.build(o<<1, l, m)
	t.build(o<<1|1, m+1, r)
}

func (t seg95) update(o, l, r, v int) {
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
	t[o].min = min(t[o<<1].min, t[o<<1|1].min)
}

func (t seg95) query(o, l, r int) int {
	if l <= t[o].l && t[o].r <= r {
		return t[o].min
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.query(o<<1, l, r)
	}
	if l > m {
		return t.query(o<<1|1, l, r)
	}
	return min(t.query(o<<1, l, r), t.query(o<<1|1, l, r))
}

func cf1295E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n int
	Fscan(in, &n)
	t := make(seg95, 2<<bits.Len(uint(n)))
	t.build(1, 0, n)
	a := make([]struct{ v, cost int }, n)
	for i := range a {
		Fscan(in, &a[i].v)
	}
	for i := range a {
		Fscan(in, &a[i].cost)
		t.update(1, a[i].v, n, a[i].cost)
	}
	ans := int(1e18)
	for _, p := range a[:n-1] {
		t.update(1, 0, p.v-1, p.cost)
		t.update(1, p.v, n, -p.cost)
		ans = min(ans, t[1].min)
	}
	Fprint(out, ans)
}

//func main() { cf1295E(os.Stdin, os.Stdout) }
