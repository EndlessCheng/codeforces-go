package main

// github.com/EndlessCheng/codeforces-go
func deleteTreeNodes(n int, parent []int, a []int) (ans int) {
	g := make([][]int, n)
	for w := 1; w < n; w++ {
		g[parent[w]] = append(g[parent[w]], w)
	}
	var f func(v, fa int) (int, int, int)
	f = func(v, fa int) (int, int, int) {
		sum, sz, del := a[v], 1, 0
		for _, w := range g[v] {
			if w != fa {
				m, s, d := f(w, v)
				sum += m
				sz += s
				del += d
			}
		}
		if sum == 0 {
			del = sz
		}
		return sum, sz, del
	}
	_, _, ans = f(0, -1)
	ans = n - ans
	return
}
