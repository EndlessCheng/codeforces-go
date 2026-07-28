package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
	"os"
	"slices"
)

// https://github.com/EndlessCheng
type seg []struct{ max, add int }

func (t seg) apply(o, f int) {
	t[o].max += f
	t[o].add += f
}

func (t seg) spread(o int) {
	f := t[o].add
	if f == 0 {
		return
	}
	t.apply(o<<1, f)
	t.apply(o<<1|1, f)
	t[o].add = 0
}

func (t seg) update(o, l, r, ql, qr, f int) {
	if ql <= l && r <= qr {
		t.apply(o, f)
		return
	}
	t.spread(o)
	m := (l + r) >> 1
	if ql <= m {
		t.update(o<<1, l, m, ql, qr, f)
	}
	if m < qr {
		t.update(o<<1|1, m+1, r, ql, qr, f)
	}
	t[o].max = max(t[o<<1].max, t[o<<1|1].max)
}

func run(in io.Reader, out io.Writer) {
	var n, d, w, ans, l int
	Fscan(in, &n, &d, &w)
	type pair struct{ t, x int }
	a := make([]pair, n)
	for i := range a {
		Fscan(in, &a[i].t, &a[i].x)
	}
	slices.SortFunc(a, func(a, b pair) int { return a.t - b.t })

	const mx = 2e5
	t := make(seg, 2<<bits.Len(mx))
	for _, p := range a {
		t.update(1, 1, mx, max(p.x-w+1, 1), p.x, 1)
		for a[l].t <= p.t-d {
			t.update(1, 1, mx, max(a[l].x-w+1, 1), a[l].x, -1)
			l++
		}
		ans = max(ans, t[1].max)
	}
	Fprint(out, ans)
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
