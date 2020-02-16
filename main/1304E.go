package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// github.com/EndlessCheng/codeforces-go
func CF1304E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, q, v, w, x, y, a, b, k int
	Fscan(in, &n)
	g := make([][]int, n)
	for i := 0; i < n-1; i++ {
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}
	vs := make([]int, 0, 2*n-1)
	pos := make([]int, n)
	depths := make([]int, 0, 2*n-1)
	dis := make([]int, n)
	var dfs func(v, fa, d int)
	dfs = func(v, fa, d int) {
		pos[v] = len(vs)
		vs = append(vs, v)
		depths = append(depths, d)
		dis[v] = d
		for _, w := range g[v] {
			if w != fa {
				dfs(w, v, d+1)
				vs = append(vs, v)
				depths = append(depths, d)
			}
		}
	}
	dfs(0, -1, 0)

	type pair struct{ v, i int }
	var st [][18]pair
	stInit := func(a []int) {
		n := len(a)
		st = make([][18]pair, n)
		for i := range st {
			st[i][0] = pair{a[i], i}
		}
		for j := uint(1); 1<<j <= n; j++ {
			for i := 0; i+(1<<j)-1 < n; i++ {
				st0, st1 := st[i][j-1], st[i+(1<<(j-1))][j-1]
				if st0.v < st1.v {
					st[i][j] = st0
				} else {
					st[i][j] = st1
				}
			}
		}
	}
	stInit(depths)
	stQuery := func(l, r int) int {
		k := uint(bits.Len(uint(r-l+1)) - 1)
		st0, st1 := st[l][k], st[r-(1<<k)+1][k]
		if st0.v < st1.v {
			return st0.i
		}
		return st1.i
	}
	lca := func(v, w int) int {
		pv, pw := pos[v], pos[w]
		if pv > pw {
			pv, pw = pw, pv
		}
		return vs[stQuery(pv, pw)]
	}
	_d := func(v, w int) int { return dis[v] + dis[w] - dis[lca(v, w)]<<1 }

_q:
	for Fscan(in, &q); q > 0; q-- {
		Fscan(in, &x, &y, &a, &b, &k)
		x--
		y--
		a--
		b--
		for _, d := range []int{_d(a, b), _d(a, x) + _d(b, y) + 1, _d(a, y) + _d(b, x) + 1} {
			if d <= k && d&1 == k&1 {
				Fprintln(out, "YES")
				continue _q
			}
		}
		Fprintln(out, "NO")
	}
}

//func main() { CF1304E(os.Stdin, os.Stdout) }
