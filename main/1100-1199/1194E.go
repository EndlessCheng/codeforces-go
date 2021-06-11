package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1194E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	const mx = 5001
	type line struct{ p, l, r int }

	var n, x1, y1, x2, y2 int
	var hs, vs []line
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &x1, &y1, &x2, &y2)
		x1 += mx
		y1 += mx
		x2 += mx
		y2 += mx
		if y1 == y2 {
			if x1 > x2 {
				x1, x2 = x2, x1
			}
			hs = append(hs, line{y1, x1, x2})
		} else {
			if y1 > y2 {
				y1, y2 = y2, y1
			}
			vs = append(vs, line{x1, y1, y2})
		}
	}
	sort.Slice(hs, func(i, j int) bool { return hs[i].p < hs[j].p })
	sort.Slice(vs, func(i, j int) bool { return vs[i].l < vs[j].l })

	ans := int64(0)
	for i, p := range hs {
		tree := [mx * 2]int{}
		add := func(i int) {
			for ; i < mx*2; i += i & -i {
				tree[i]++
			}
		}
		sum := func(i int) (res int) {
			for ; i > 0; i &= i - 1 {
				res += tree[i]
			}
			return
		}
		j := 0
		for _, q := range hs[:i] {
			for ; j < len(vs) && vs[j].l <= q.p; j++ {
				if vs[j].r >= p.p {
					add(vs[j].p)
				}
			}
			if l, r := max(q.l, p.l), min(q.r, p.r); l <= r {
				c := sum(r) - sum(l-1)
				ans += int64(c * (c - 1) / 2)
			}
		}
	}
	Fprint(out, ans)
}

//func main() { CF1194E(os.Stdin, os.Stdout) }
