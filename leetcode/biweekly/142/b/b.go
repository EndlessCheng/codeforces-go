package main

// https://space.bilibili.com/206214
func findSubtreeSizes(parent []int, s string) []int {
	n := len(parent)
	g := make([][]int, n)
	for i := 1; i < n; i++ {
		p := parent[i]
		g[p] = append(g[p], i)
	}

	size := make([]int, n)
	ancestor := [26]int{}
	for i := range ancestor {
		ancestor[i] = -1
	}
	var dfs func(int)
	dfs = func(x int) {
		size[x] = 1
		sx := s[x] - 'a'
		old := ancestor[sx]
		ancestor[sx] = x
		for _, y := range g[x] {
			dfs(y)
			anc := ancestor[s[y]-'a']
			if anc < 0 {
				anc = x
			}
			size[anc] += size[y]
		}
		ancestor[sx] = old // 恢复现场
	}
	dfs(0)
	return size
}

func findSubtreeSizes2(parent []int, s string) []int {
	n := len(parent)
	g := make([][]int, n)
	for i := 1; i < n; i++ {
		p := parent[i]
		g[p] = append(g[p], i)
	}

	ancestor := [26]int{}
	for i := range ancestor {
		ancestor[i] = -1
	}
	var rebuild func(int)
	rebuild = func(x int) {
		sx := s[x] - 'a'
		old := ancestor[sx]
		ancestor[sx] = x
		for i, y := range g[x] {
			if anc := ancestor[s[y]-'a']; anc != -1 {
				g[anc] = append(g[anc], y)
				g[x][i] = -1 // 删除 y
			}
			rebuild(y)
		}
		ancestor[sx] = old // 恢复现场
	}
	rebuild(0)

	size := make([]int, n)
	var dfs func(int)
	dfs = func(x int) {
		size[x] = 1
		for _, y := range g[x] {
			if y != -1 {
				dfs(y)
				size[x] += size[y]
			}
		}
	}
	dfs(0)
	return size
}
