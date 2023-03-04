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
	has := make(map[pair]bool, len(guesses))
	for _, p := range guesses { // guesses 转成哈希表
		has[pair{p[0], p[1]}] = true
	}

	cnt0 := 0
	var dfs func(int, int)
	dfs = func(x, fa int) {
		for _, y := range g[x] {
			if y != fa {
				if has[pair{x, y}] { // 以 0 为根时，猜对了
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
				c := cnt
				if has[pair{x, y}] { // 原来是对的，现在错了
					c--
				}
				if has[pair{y, x}] { // 原来是错的，现在对了
					c++
				}
				reroot(y, x, c)
			}
		}
	}
	reroot(0, -1, cnt0)
	return
}
