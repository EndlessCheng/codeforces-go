package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// https://space.bilibili.com/206214
type seg42 []struct{ l, r, max, todo int }

func (t seg42) do(o, v int) {
	t[o].max += v
	t[o].todo += v
}

func (t seg42) spread(o int) {
	if v := t[o].todo; v != 0 {
		t.do(o<<1, v)
		t.do(o<<1|1, v)
		t[o].todo = 0
	}
}

func (t seg42) build(o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		return
	}
	m := (l + r) >> 1
	t.build(o<<1, l, m)
	t.build(o<<1|1, m+1, r)
}

func (t seg42) update(o, l, r int, v int) {
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
	t[o].max = max(t[o<<1].max, t[o<<1|1].max)
}

func (t seg42) query(o, l, r int) int {
	if l <= t[o].l && t[o].r <= r {
		return t[o].max
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

func cf1842E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, k, A, sc, f int
	Fscan(in, &n, &k, &A)
	type pair struct{ x, c int }
	ps := make([][]pair, k+1)
	for i := 0; i < n; i++ {
		var x, y, c int
		Fscan(in, &x, &y, &c)
		sc += c
		ps[k-y] = append(ps[k-y], pair{x, c})
	}

	t := make(seg42, 2<<bits.Len(uint(k)))
	t.build(1, 0, k)
	for i := 1; i <= k; i++ {
		t.update(1, 0, i-1, -A)
		for _, p := range ps[i] {
			t.update(1, 0, p.x, p.c)
		}
		f = max(f, t[1].max)
		t.update(1, i, i, f)
	}
	Fprint(out, sc-f)
}

//func main() { cf1842E(os.Stdin, os.Stdout) }
