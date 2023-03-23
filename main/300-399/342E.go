package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
	"math/bits"
)

// https://space.bilibili.com/206214
func CF342E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}

	var n, m, op, v, w int
	Fscan(in, &n, &m)
	g := make([][]int, n)
	for i := 1; i < n; i++ {
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	vs := make([]int, 0, 2*n-1)
	pos := make([]int, n)
	dep := make([]int, 0, 2*n-1)
	disRoot := make([]int, n)
	var build func(v, p, d int)
	build = func(v, p, d int) {
		pos[v] = len(vs)
		vs = append(vs, v)
		dep = append(dep, d)
		disRoot[v] = d
		for _, w := range g[v] {
			if w != p {
				build(w, v, d+1)
				vs = append(vs, v)
				dep = append(dep, d)
			}
		}
	}
	build(1, -1, 0)
	type stPair struct{ v, i int }
	const mx = 18
	var st [][mx]stPair
	stInit := func(a []int) {
		n := len(a)
		st = make([][mx]stPair, n)
		for i, v := range a {
			st[i][0] = stPair{v, i}
		}
		for j := 1; 1<<j <= n; j++ {
			for i := 0; i+1<<j <= n; i++ {
				if a, b := st[i][j-1], st[i+1<<(j-1)][j-1]; a.v < b.v {
					st[i][j] = a
				} else {
					st[i][j] = b
				}
			}
		}
	}
	stInit(dep)
	stQuery := func(l, r int) int {
		k := bits.Len(uint(r-l)) - 1
		a, b := st[l][k], st[r-1<<k][k]
		if a.v < b.v {
			return a.i
		}
		return b.i
	}
	lca := func(v, w int) int {
		pv, pw := pos[v], pos[w]
		if pv > pw {
			pv, pw = pw, pv
		}
		return vs[stQuery(pv, pw+1)]
	}
	_d := func(v, w int) int { return disRoot[v] + disRoot[w] - disRoot[lca(v, w)]<<1 }

	dis := make([]int, n)
	for i := range dis {
		dis[i] = 1e9
	}
	type pair struct{ v, fa int }
	q0 := make([]pair, 0, n)
	bfs := func(q []pair) {
		for _, p := range q {
			dis[p.v] = 0
		}
		for len(q) > 0 {
			p := q[0]
			q = q[1:]
			v := p.v
			for _, w := range g[v] {
				if w != p.fa && dis[v]+1 < dis[w] {
					dis[w] = dis[v] + 1
					q = append(q, pair{w, v})
				}
			}
		}
	}
	bfs(append(q0, pair{0, -1}))

	sqSize := int(math.Sqrt(float64(m)))
	for q := q0; m > 0; m-- {
		Fscan(in, &op, &v)
		v--
		if op == 1 {
			q = append(q, pair{v, -1})
			if len(q) == sqSize {
				bfs(q)
				q = q0
			}
		} else {
			ans := dis[v]
			for _, p := range q {
				ans = min(ans, _d(v, p.v))
			}
			Fprintln(out, ans)
		}
	}
}

//func main() { CF342E(os.Stdin, os.Stdout) }
