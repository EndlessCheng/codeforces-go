package main

// github.com/EndlessCheng/codeforces-go
func longestPath(parent []int, s string) (ans int) {
	n := len(parent)
	g := make([][]int, n)
	for i := 1; i < n; i++ {
		pa := parent[i]
		g[pa] = append(g[pa], i)
	}

	var dfs func(int) int
	dfs = func(x int) (maxLen int) {
		for _, y := range g[x] {
			len := dfs(y) + 1
			if s[y] != s[x] {
				ans = max(ans, len+maxLen)
				maxLen = max(maxLen, len)
			}
		}
		return
	}
	dfs(0)
	return ans + 1
}

func max(a, b int) int { if b > a { return b }; return a }
