package main

import "strconv"

// github.com/EndlessCheng/codeforces-go
func digSum(a []int, connectRoad [][]int) string {
	n := len(a)
	g := make([][]int, n)
	deg := make([]int, n)
	for _, e := range connectRoad {
		v, w := e[0]-1, e[1]-1
		if v > w {
			v, w = w, v
		}
		g[v] = append(g[v], w)
		deg[w]++
	}

	fa := make([]int, n)
	for i := range fa {
		fa[i] = -1
	}
	dp := make([]int, n)
	q := []int{}
	for i, d := range deg {
		if d == 0 {
			dp[i] = a[i]
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		for _, w := range g[v] {
			if dp[v]+a[w] > dp[w] {
				dp[w] = dp[v] + a[w]
				fa[w] = v
			}
			if deg[w]--; deg[w] == 0 {
				q = append(q, w)
			}
		}
	}

	end := 0
	for i, v := range dp {
		if v > dp[end] {
			end = i
		}
	}
	path := make([]int, 0, n)
	for v := end; v != -1; v = fa[v] {
		path = append(path, v)
	}
	ans := []byte{}
	for i := len(path) - 1; i >= 0; i-- {
		if i < len(path)-1 {
			ans = append(ans, '-')
		}
		ans = append(ans, strconv.Itoa(path[i]+1)...)
	}
	return string(ans)
}