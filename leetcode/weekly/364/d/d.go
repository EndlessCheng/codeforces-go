package main

// https://space.bilibili.com/206214
const mx int = 1e5 + 1

var np = [mx]bool{1: true}

func init() { // 质数=false 非质数=true
	for i := 2; i*i < mx; i++ {
		if !np[i] {
			for j := i * i; j < mx; j += i {
				np[j] = true
			}
		}
	}
}

func countPaths(n int, edges [][]int) (ans int64) {
	g := make([][]int, n+1)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	size := make([]int, n+1)
	nodes := []int{}
	var dfs func(int, int)
	dfs = func(x, fa int) {
		nodes = append(nodes, x)
		for _, y := range g[x] {
			if y != fa && np[y] {
				dfs(y, x)
			}
		}
	}
	for x := 1; x <= n; x++ {
		if np[x] { // 跳过非质数
			continue
		}
		sum := 1
		for _, y := range g[x] { // 质数 x 把这棵树分成了若干个连通块
			if !np[y] {
				continue
			}
			if size[y] == 0 {
				nodes = nodes[:0]
				dfs(y, -1) // 遍历 y 所在连通块，在不经过质数的前提下，统计有多少个非质数
				for _, z := range nodes {
					size[z] = len(nodes)
				}
			}
			// 这 cnt 个非质数与之前遍历到的 sum 个非质数，两两之间的路径只包含质数 x
			ans += int64(size[y]) * int64(sum)
			sum += size[y]
		}
	}
	return
}
