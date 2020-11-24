package main

// github.com/EndlessCheng/codeforces-go
func tree6(k int, a []int) int64 {
	n := len(a)
	g := make([][]int, n)
	for w := 1; w < n; w++ {
		v := (w - 1) / k
		g[v] = append(g[v], w)
	}
	ans, p := 0, 0
	var f func(int)
	f = func(v int) {
		val := a[p]
		for _, w := range g[v] {
			p++
			ans += val ^ a[p]
			f(w)
		}
	}
	f(0)
	return int64(ans)
}
