package main

// github.com/EndlessCheng/codeforces-go
func numFlowers(es [][]int) (ans int) {
	g := make([][]int, len(es)+1)
	for _, e := range es {
		v, w := e[0], e[1]
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}
	for _, vs := range g {
		ans = max(ans, len(vs))
	}
	ans++
	return
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
