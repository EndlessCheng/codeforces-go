package main

// https://space.bilibili.com/206214
func longestString(x, y, z int) int {
	memo := make([][][][3]int, x+1)
	for i := range memo {
		memo[i] = make([][][3]int, y+1)
		for j := range memo[i] {
			memo[i][j] = make([][3]int, z+1)
			for k := range memo[i][j] {
				memo[i][j][k] = [3]int{-1, -1, -1}
			}
		}
	}
	var dfs func(x, y, z, k int) int
	dfs = func(x, y, z, k int) (res int) {
		p := &memo[x][y][z][k]
		if *p != -1 { // 之前算过
			return *p
		}
		if k == 0 {
			if y > 0 {
				res = dfs(x, y-1, z, 1) + 2
			}
		} else {
			if x > 0 {
				res = dfs(x-1, y, z, 0) + 2
			}
			if z > 0 {
				res = max(res, dfs(x, y, z-1, 2)+2)
			}
		}
		*p = res // 记忆化
		return
	}
	return max(dfs(x, y, z, 0), dfs(x, y, z, 1))
}

func max(a, b int) int { if b > a { return b }; return a }
