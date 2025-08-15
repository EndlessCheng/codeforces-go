package main

import (
	. "fmt"
	"io"
	"math/bits"
)

// https://github.com/EndlessCheng
type seg15 []struct{ l, r, mx, todo int }

func (t seg15) apply(o, f int) {
	t[o].mx += f
	t[o].todo += f
}

func (t seg15) maintain(o int) {
	t[o].mx = max(t[o<<1].mx, t[o<<1|1].mx)
}

func (t seg15) spread(o int) {
	f := t[o].todo
	if f == 0 {
		return
	}
	t.apply(o<<1, f)
	t.apply(o<<1|1, f)
	t[o].todo = 0
}

func (t seg15) build(o, l, r int) {
	t[o].l, t[o].r = l, r
	t[o].mx = -1e18
	if l == r {
		return
	}
	m := (l + r) >> 1
	t.build(o<<1, l, m)
	t.build(o<<1|1, m+1, r)
}

func (t seg15) set(o, i, v int) {
	if t[o].l == t[o].r {
		t[o].mx = v
		return
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if i <= m {
		t.set(o<<1, i, v)
	} else {
		t.set(o<<1|1, i, v)
	}
	t.maintain(o)
}

func (t seg15) update(o, l, r, f int) {
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

func (t seg15) query(o, l, r int) int {
	if l <= t[o].l && t[o].r <= r {
		return t[o].mx
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.query(o<<1, l, r)
	}
	if l > m {
		return t.query(o<<1|1, l, r)
	}
	return max(t.query(o<<1, l, r), t.query(o<<1|1, l, r))
}

func cf115E(in io.Reader, out io.Writer) {
	var n, m, l, r, p, f int
	Fscan(in, &n, &m)
	s := make([]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &s[i])
		s[i] += s[i-1]
	}
	type pair struct{ l, p int }
	g := make([][]pair, n+1)
	for range m {
		Fscan(in, &l, &r, &p)
		g[r] = append(g[r], pair{l, p})
	}

	t := make(seg15, 2<<bits.Len(uint(n-1)))
	t.build(1, 1, n)
	for i := 1; i <= n; i++ {
		t.set(1, i, f+s[i-1])
		for _, p := range g[i] {
			t.update(1, 1, p.l, p.p)
		}
		f = max(f, t.query(1, 1, i)-s[i])
	}
	Fprint(out, f)
}

//func main() { cf115E(bufio.NewReader(os.Stdin), os.Stdout) }
