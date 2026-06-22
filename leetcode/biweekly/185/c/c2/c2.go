package main

import "math"

// https://space.bilibili.com/206214
func finishTime(n int, edges [][]int, baseTime []int) int64 {
	g := make([][]int, n)
	for _, e := range edges {
		v, w := e[0], e[1]
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	subRes := make([]struct{ mx, mx2, mxW, mn, mn2, mnW int }, n)
	var dfs func(int, int) int
	dfs = func(v, fa int) int {
		res := &subRes[v]
		res.mn = math.MaxInt
		res.mn2 = math.MaxInt
		for _, w := range g[v] {
			if w == fa {
				continue
			}
			r := dfs(w, v)
			if r > res.mx {
				res.mx2 = res.mx
				res.mx = r
				res.mxW = w
			} else if r > res.mx2 {
				res.mx2 = r
			}
			if r < res.mn {
				res.mn2 = res.mn
				res.mn = r
				res.mnW = w
			} else if r < res.mn2 {
				res.mn2 = r
			}
		}
		if res.mx == 0 {
			return baseTime[v]
		}
		return res.mx*2 - res.mn + baseTime[v]
	}
	ans := dfs(0, -1)

	var reroot func(int, int, int)
	reroot = func(v, fa, fromUp int) {
		res := subRes[v]
		if fromUp > 0 {
			ans = min(ans, max(res.mx, fromUp)*2-min(res.mn, fromUp)+baseTime[v])
		}
		for _, w := range g[v] {
			if w == fa {
				continue
			}
			t := baseTime[v]
			if res.mx2 > 0 {
				mx, mn := res.mx, res.mn
				if w == res.mxW {
					mx = res.mx2
				} else if w == res.mnW {
					mn = res.mn2
				}
				if fromUp > 0 {
					mx = max(mx, fromUp)
					mn = min(mn, fromUp)
				}
				t += mx*2 - mn
			} else {
				t += fromUp
			}
			reroot(w, v, t)
		}
	}
	reroot(0, -1, 0)
	return int64(ans)
}
