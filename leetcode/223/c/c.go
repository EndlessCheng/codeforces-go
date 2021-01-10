package main

// github.com/EndlessCheng/codeforces-go
func minimumHammingDistance(a, b []int, es [][]int) (ans int) {
	n := len(a)
	g := make([][]int, n)
	for _, e := range es {
		v, w := e[0], e[1]
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}
	vis := make([]bool, n)
	var c1, c2 map[int]int
	var f func(int)
	f = func(v int) {
		vis[v] = true
		c1[a[v]]++
		c2[b[v]]++
		for _, w := range g[v] {
			if !vis[w] {
				f(w)
			}
		}
	}
	for i, b := range vis {
		if !b {
			c1, c2 = map[int]int{}, map[int]int{}
			f(i)
			for k, c := range c1 {
				ans += min(c, c2[k])
			}
		}
	}
	return n - ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
