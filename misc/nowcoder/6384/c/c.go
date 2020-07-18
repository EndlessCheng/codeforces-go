package main

import "strconv"

// github.com/EndlessCheng/codeforces-go
func digSum(a []int, connectRoad [][]int) string {
	n := len(a)
	g := make([][]int, n+1)
	inDeg := make([]int, n+1)
	for _, e := range connectRoad {
		v, w := e[0], e[1]
		if v > w {
			v, w = w, v
		}
		g[v] = append(g[v], w)
		inDeg[w]++
	}

	fa := make([]int, n+1)
	for i := range fa {
		fa[i] = -1
	}
	q := []int{}
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		if inDeg[i] == 0 {
			q = append(q, i)
			dp[i] = a[i-1]
		}
	}
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		for _, w := range g[v] {
			inDeg[w]--
			if dp[v]+a[w-1] > dp[w] {
				dp[w] = dp[v] + a[w-1]
				fa[w] = v
			}
			if inDeg[w] == 0 {
				q = append(q, w)
			}
		}
	}

	end := 1
	for i := 2; i <= n; i++ {
		if dp[i] > dp[end] {
			end = i
		}
	}
	path := make([]int, 0, n)
	for x := end; x != -1; x = fa[x] {
		path = append(path, x)
	}
	ans := []byte{}
	for i := len(path) - 1; i >= 0; i-- {
		v := path[i]
		if i < len(path)-1 {
			ans = append(ans, '-')
		}
		ans = append(ans, strconv.Itoa(v)...)
	}
	return string(ans)
}
