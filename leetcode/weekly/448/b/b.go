package main

// https://space.bilibili.com/206214
func specialGrid(n int) [][]int {
	val := 0
	var dfs func([][]int, int)
	dfs = func(a [][]int, l int) {
		if len(a) == 1 {
			a[0][l] = val
			val++
			return
		}
		m := len(a) / 2
		dfs(a[:m], l+m)
		dfs(a[m:], l+m)
		dfs(a[m:], l)
		dfs(a[:m], l)
	}

	a := make([][]int, 1<<n)
	for i := range a {
		a[i] = make([]int, 1<<n)
	}
	dfs(a, 0)
	return a
}

func specialGrid2(n int) [][]int {
	a := make([][]int, 1<<n)
	for i := range a {
		a[i] = make([]int, 1<<n)
	}
	val := 0
	var dfs func(int, int, int, int)
	dfs = func(u, d, l, r int) {
		if d-u == 1 {
			a[u][l] = val
			val++
			return
		}
		m := (d - u) / 2
		dfs(u, u+m, l+m, r)
		dfs(u+m, d, l+m, r)
		dfs(u+m, d, l, l+m)
		dfs(u, u+m, l, l+m)
	}
	dfs(0, 1<<n, 0, 1<<n)
	return a
}
