package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
	"os"
)

// https://github.com/EndlessCheng
type seg []struct {
	minCover    int
	minCoverSum int
	add         int
}

func (t seg) apply(o, x int) {
	t[o].minCover += x
	t[o].add += x
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

func (t seg) maintain(o int) {
	cur, lo, ro := &t[o], &t[o<<1], &t[o<<1|1]
	mn := min(lo.minCover, ro.minCover)
	cur.minCover = mn
	cur.minCoverSum = 0
	if lo.minCover == mn {
		cur.minCoverSum = lo.minCoverSum
	}
	if ro.minCover == mn {
		cur.minCoverSum += ro.minCoverSum
	}
}

func (t seg) build(o, l, r int) {
	if l == r {
		t[o].minCoverSum = l
		return
	}
	m := (l + r) >> 1
	t.build(o<<1, l, m)
	t.build(o<<1|1, m+1, r)
	t.maintain(o)
}

func (t seg) update(o, l, r, ql, qr, x int) {
	if ql <= l && r <= qr {
		t.apply(o, x)
		return
	}
	t.spread(o)
	m := (l + r) >> 1
	if ql <= m {
		t.update(o<<1, l, m, ql, qr, x)
	}
	if m < qr {
		t.update(o<<1|1, m+1, r, ql, qr, x)
	}
	t.maintain(o)
}

func run(in io.Reader, out io.Writer) {
	const mod = 998244353
	var n, q, r1, c1, r2, c2, ans int
	Fscan(in, &n, &q)
	type tuple struct{ l, r, add int }
	rows := make([][]tuple, n+2)
	for range q {
		Fscan(in, &r1, &c1, &r2, &c2)
		rows[r1] = append(rows[r1], tuple{c1, c2, 1})
		rows[r2+1] = append(rows[r2+1], tuple{c1, c2, -1})
	}

	t := make(seg, 2<<bits.Len(uint(n-1)))
	t.build(1, 1, n)
	for i := 1; i <= n; i++ {
		for _, p := range rows[i] {
			t.update(1, 1, n, p.l, p.r, p.add)
		}
		if t[1].minCover == 0 {
			ans = (ans + i*t[1].minCoverSum) % mod
		}
	}
	Fprint(out, ans)
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
