package main

// https://space.bilibili.com/206214
func rootCount(edges [][]int, guesses [][]int, k int) (ans int) {
	g := make([][]int, len(edges)+1)
	for _, e := range edges {
		v, w := e[0], e[1]
		g[v] = append(g[v], w)
		g[w] = append(g[w], v) // 建图
	}

	type pair struct{ x, y int }
	s := make(map[pair]int, len(guesses))
	for _, p := range guesses { // guesses 转成哈希表
		s[pair{p[0], p[1]}] = 1
	}

	cnt0 := 0
	var dfs func(int, int)
	dfs = func(x, fa int) {
		for _, y := range g[x] {
			if y != fa {
				if s[pair{x, y}] == 1 { // 以 0 为根时，猜对了
					cnt0++
				}
				dfs(y, x)
			}
		}
	}
	dfs(0, -1)

	var reroot func(int, int, int)
	reroot = func(x, fa, cnt int) {
		if cnt >= k { // 此时 cnt 就是以 x 为根时的猜对次数
			ans++
		}
		for _, y := range g[x] {
			if y != fa {
				reroot(y, x, cnt-s[pair{x, y}]+s[pair{y, x}])
			}
		}
	}
	reroot(0, -1, cnt0)
	return
}
