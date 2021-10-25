package main

// github.com/EndlessCheng/codeforces-go
func countHighestScoreNodes(parents []int) (ans int) {
	n := len(parents)
	g := make([][]int, n)
	for w := 1; w < n; w++ {
		v := parents[w]
		g[v] = append(g[v], w) // 建树
	}

	maxScore := 0
	var dfs func(int) int
	dfs = func(v int) int {
		size, score := 1, 1
		for _, w := range g[v] {
			sz := dfs(w)
			size += sz
			score *= sz // 由于是二叉树所以 score 最大约为 (1e5/3)^3，在 64 位整数范围内
		}
		if v > 0 {
			score *= n - size
		}
		if score > maxScore {
			maxScore, ans = score, 1
		} else if score == maxScore {
			ans++
		}
		return size
	}
	dfs(0)
	return
}
