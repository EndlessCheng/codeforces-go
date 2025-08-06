package main

import (
	. "fmt"
	"io"
	"math/bits"
)

// https://github.com/EndlessCheng
type seg69 []struct{ l, r, mn, todo int }

func (t seg69) apply(o, f int) {
	t[o].mn += f
	t[o].todo += f
}

func (t seg69) spread(o int) {
	f := t[o].todo
	if f == 0 {
		return
	}
	t.apply(o<<1, f)
	t.apply(o<<1|1, f)
	t[o].todo = 0
}

func (t seg69) build(o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		return
	}
	m := (l + r) >> 1
	t.build(o<<1, l, m)
	t.build(o<<1|1, m+1, r)
}

func (t seg69) update(o, l, r, f int) {
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
	t[o].mn = min(t[o<<1].mn, t[o<<1|1].mn)
}

func (t seg69) query(o, l, r int) int {
	if l <= t[o].l && t[o].r <= r {
		return t[o].mn
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

func cf1969E(in io.Reader, out io.Writer) {
	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		t := make(seg69, 2<<bits.Len(uint(n-1)))
		t.build(1, 1, n)
		pre := make([]int, n+1)
		pre2 := make([]int, n+1)

		ans, l := 0, 1
		for i := 1; i <= n; i++ {
			Fscan(in, &v)
			if pre[v] > 0 {
				t.update(1, pre2[v]+1, pre[v], -1)
			}
			t.update(1, pre[v]+1, i, 1)
			pre2[v] = pre[v]
			pre[v] = i
			if t.query(1, l, i) == 0 {
				ans++
				l = i + 1
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1969E(bufio.NewReader(os.Stdin), os.Stdout) }
