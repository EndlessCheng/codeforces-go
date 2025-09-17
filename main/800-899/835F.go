package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf835F(in io.Reader, out io.Writer) {
	var n, diam, maxInnerDiam int
	Fscan(in, &n)
	type nb struct{ to, wt int }
	g := make([][]nb, n)
	deg := make([]int, n)
	for range n {
		var v, w, wt int
		Fscan(in, &v, &w, &wt)
		v--
		w--
		g[v] = append(g[v], nb{w, wt})
		g[w] = append(g[w], nb{v, wt})
		deg[v]++
		deg[w]++
	}

	q := []int{}
	for i, d := range deg {
		if d == 1 {
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		for _, e := range g[v] {
			w := e.to
			deg[w]--
			if deg[w] == 1 {
				q = append(q, w)
			}
		}
	}

	cycle := []nb{}
	for start, d := range deg {
		if d > 1 {
			pre := -1
			cur := start
			for {
				cycle = append(cycle, nb{cur, 0})
				for _, e := range g[cur] {
					w := e.to
					if w != pre && deg[w] > 1 {
						cycle[len(cycle)-1].wt = e.wt
						pre = cur
						cur = w
						break
					}
				}
				if cur == start {
					break
				}
			}
			break
		}
	}
	m := len(cycle)

	var dfs func(int, int) int
	dfs = func(v, fa int) (maxL int) {
		for _, e := range g[v] {
			w := e.to
			if w != fa && deg[w] < 2 {
				subL := dfs(w, v) + e.wt
				diam = max(diam, maxL+subL)
				maxL = max(maxL, subL)
			}
		}
		return
	}
	maxL := make([]int, m)
	for i, e := range cycle {
		diam = 0
		maxL[i] = dfs(e.to, -1)
		maxInnerDiam = max(maxInnerDiam, diam)
	}

	suf := make([]int, m+1)
	suf2 := make([]int, m)
	s, mx := 0, maxL[m-1]
	for i := m - 2; i >= 0; i-- {
		suf[i+1] = max(suf[i+2], maxL[i+1]+s)
		s += cycle[i].wt
		suf2[i] = max(suf2[i+1], maxL[i]+s+mx)
		mx = max(mx, maxL[i]-s)
	}
	ans := suf2[0]

	pre, pre2, s, mx := 0, 0, 0, maxL[0]
	for i := range m - 1 {
		pre = max(pre, maxL[i]+s)
		ans = min(ans, max(pre2, suf2[i+1], pre+cycle[m-1].wt+suf[i+1]))
		s += cycle[i].wt
		pre2 = max(pre2, maxL[i+1]+s+mx)
		mx = max(mx, maxL[i+1]-s)
	}

	Fprint(out, max(ans, maxInnerDiam))
}

//func main() { cf835F(bufio.NewReader(os.Stdin), os.Stdout) }
