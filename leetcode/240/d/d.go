package main

// github.com/EndlessCheng/codeforces-go
func largestPathValue(s string, es [][]int) (ans int) {
	n := len(s)
	g := make([][]int, n)
	deg := make([]int, n)
	for _, e := range es {
		v, w := e[0], e[1]
		if v == w {
			return -1
		}
		g[v] = append(g[v], w)
		deg[w]++
	}
	orders := []int{}
	dp := make([][26]int, n)
	q := []int{}
	for i, d := range deg {
		if d == 0 {
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		orders = append(orders, v)
		dp[v][s[v]-'a']++
		ans = max(ans, dp[v][s[v]-'a'])
		for _, w := range g[v] {
			for i, c := range dp[v] {
				dp[w][i] = max(dp[w][i], c)
			}
			if deg[w]--; deg[w] == 0 {
				q = append(q, w)
			}
		}
	}
	if len(orders) < n {
		return -1
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
