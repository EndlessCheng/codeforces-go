package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
	"slices"
)

// https://github.com/EndlessCheng
func cf601E(in io.Reader, _w io.Writer) {
	const mod = 1_000_000_007
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, k, q, t, op, v, w int
	Fscan(in, &n, &k)
	type pair struct{ v, w int }
	a := make([]pair, n)
	type lr struct{ l, r int }
	ranges := make([]lr, n)
	for i := range a {
		Fscan(in, &a[i].v, &a[i].w)
		ranges[i].r = -1
	}

	Fscan(in, &q)
	for range q {
		Fscan(in, &op)
		if op == 1 {
			Fscan(in, &v, &w)
			a = append(a, pair{v, w})
			ranges = append(ranges, lr{t, -1})
		} else if op == 2 {
			Fscan(in, &v)
			ranges[v-1].r = t
		} else {
			t++
		}
	}

	g := make([][]pair, 2<<bits.Len(uint(t-1))) // 题目保证 t > 0
	var update func(o, l, r, ql, qr int, p pair)
	update = func(o, l, r, ql, qr int, p pair) {
		if ql <= l && r <= qr {
			g[o] = append(g[o], p)
			return
		}
		m := (l + r) / 2
		if ql <= m {
			update(o*2, l, m, ql, qr, p)
		}
		if m < qr {
			update(o*2+1, m+1, r, ql, qr, p)
		}
	}
	for i, p := range ranges {
		if p.r < 0 {
			p.r = t
		}
		if p.l < p.r {
			update(1, 0, t-1, p.l, p.r-1, a[i])
		}
	}

	var dfs func(o, l, r int, f []int)
	dfs = func(o, l, r int, f []int) {
		if g[o] != nil {
			f = slices.Clone(f)
			for _, p := range g[o] {
				for i := k; i >= p.w; i-- {
					f[i] = max(f[i], f[i-p.w]+p.v)
				}
			}
		}
		if l == r {
			ans, powP := 0, 1
			for _, v := range f[1:] {
				ans = (ans + v*powP) % mod
				powP = powP * 10_000_019 % mod
			}
			Fprintln(out, ans)
			return
		}
		m := (l + r) / 2
		dfs(o*2, l, m, f)
		dfs(o*2+1, m+1, r, f)
	}
	dfs(1, 0, t-1, make([]int, k+1))
}

//func main() { cf601E(bufio.NewReader(os.Stdin), os.Stdout) }
