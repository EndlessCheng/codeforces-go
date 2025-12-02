package main

import (
	. "fmt"
	"io"
	"slices"
	"sort"
)

// https://github.com/EndlessCheng
type node80 struct {
	lo, ro *node80
	minL   int
}

func build80(l, r int) *node80 {
	o := &node80{}
	if l == r {
		return o
	}
	m := (l + r) >> 1
	o.lo = build80(l, m)
	o.ro = build80(m+1, r)
	return o
}

func (o node80) update(l, r, i, val int) *node80 {
	if l == r {
		o.minL = max(o.minL, val)
		return &o
	}
	m := (l + r) >> 1
	if i <= m {
		o.lo = o.lo.update(l, m, i, val)
	} else {
		o.ro = o.ro.update(m+1, r, i, val)
	}
	o.minL = min(o.lo.minL, o.ro.minL)
	return &o
}

func (o *node80) query(l, r, ql, qr int) int {
	if ql <= l && r <= qr {
		return o.minL
	}
	m := (l + r) >> 1
	if qr <= m {
		return o.lo.query(l, m, ql, qr)
	}
	if m < ql {
		return o.ro.query(m+1, r, ql, qr)
	}
	return min(o.lo.query(l, m, ql, qr), o.ro.query(m+1, r, ql, qr))
}

func cf1080F(in io.Reader, out io.Writer) {
	var n, m, k, l, r, p, mn, mx int
	Fscan(in, &n, &m, &k)
	type pair struct{ l, p int }
	g := map[int][]pair{}
	for range k {
		Fscan(in, &l, &r, &p)
		g[r] = append(g[r], pair{l, p})
	}

	rs := make([]int, 0, len(g))
	for r := range g {
		rs = append(rs, r)
	}
	slices.Sort(rs)

	t := make([]*node80, len(rs)+1)
	t[0] = build80(1, n)
	for i, r := range rs {
		rt := t[i]
		for _, p := range g[r] {
			rt = rt.update(1, n, p.p, p.l)
		}
		t[i+1] = rt
	}

	for range m {
		Fscan(in, &l, &r, &mn, &mx)
		i := sort.SearchInts(rs, mx+1)
		if t[i].query(1, n, l, r) >= mn {
			Fprintln(out, "yes")
		} else {
			Fprintln(out, "no")
		}
	}
}

//func main() { debug.SetGCPercent(-1); cf1080F(bufio.NewReader(os.Stdin), os.Stdout) }
