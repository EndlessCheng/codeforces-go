package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
	"os"
)

// https://github.com/EndlessCheng
type seg []struct{ min, minCnt, todo int }

func (t seg) maintain(o int) {
	cur, lo, ro := &t[o], &t[o<<1], &t[o<<1|1]
	mn := min(lo.min, ro.min)
	cur.min = mn
	cur.minCnt = 0
	if lo.min == mn {
		cur.minCnt = lo.minCnt
	}
	if ro.min == mn {
		cur.minCnt += ro.minCnt
	}
}

func (t seg) apply(o, f int) {
	t[o].min += f
	t[o].todo += f
}

func (t seg) build(o, l, r int) {
	t[o].minCnt = r - l + 1
	if l == r {
		return
	}
	m := (l + r) >> 1
	t.build(o<<1, l, m)
	t.build(o<<1|1, m+1, r)
}

func (t seg) spread(o int) {
	f := t[o].todo
	if f == 0 {
		return
	}
	t.apply(o<<1, f)
	t.apply(o<<1|1, f)
	t[o].todo = 0
}

func (t seg) update(o, l, r, ql, qr, v int) {
	if ql <= l && r <= qr {
		t.apply(o, v)
		return
	}
	t.spread(o)
	m := (l + r) >> 1
	if ql <= m {
		t.update(o<<1, l, m, ql, qr, v)
	}
	if m < qr {
		t.update(o<<1|1, m+1, r, ql, qr, v)
	}
	t.maintain(o)
}

func (t seg) query(o, l, r, qr int) int {
	if r <= qr {
		if t[o].min == 0 {
			return t[o].minCnt
		}
		return 0
	}
	t.spread(o)
	m := (l + r) >> 1
	if qr <= m {
		return t.query(o<<1, l, m, qr)
	}
	return t.query(o<<1, l, m, qr) + t.query(o<<1|1, m+1, r, qr)
}

func run(in io.Reader, out io.Writer) {
	var n, v, ans int
	Fscan(in, &n)
	last := make([]int, n+1)
	last2 := make([]int, n+1)
	t := make(seg, 2<<bits.Len(uint(n-1)))
	t.build(1, 1, n)
	for i := 1; i <= n; i++ {
		Fscan(in, &v)
		j := last[v]
		if j > 0 {
			t.update(1, 1, n, last2[v]+1, j, -1)
		}
		t.update(1, 1, n, j+1, i, 1)
		ans += i - t.query(1, 1, n, i)
		last2[v] = j
		last[v] = i
	}
	Fprint(out, ans)
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
