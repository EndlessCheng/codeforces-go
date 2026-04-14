package main

// github.com/EndlessCheng/codeforces-go
func minimumHammingDistance(source []int, target []int, allowedSwaps [][]int) (ans int) {
	n := len(source)
	g := make([][]int, n)
	for _, e := range allowedSwaps {
		i, j := e[0], e[1]
		g[i] = append(g[i], j) // 建图
		g[j] = append(g[j], i)
	}

	vis := make([]bool, n)
	diff := map[int]int{}

	var dfs func(int)
	dfs = func(x int) {
		vis[x] = true // 避免重复访问
		// 抵消相同的元素，最终剩下 source 和 target 各自多出来的元素（对称差）
		diff[source[x]]++
		diff[target[x]]--
		for _, y := range g[x] {
			if !vis[y] {
				dfs(y)
			}
		}
	}

	for x, b := range vis {
		if !b {
			diff = map[int]int{}
			dfs(x)
			for _, c := range diff {
				ans += abs(c)
			}
		}
	}
	return ans / 2 // 有 ans / 2 对多出来的元素
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
