package main

// https://space.bilibili.com/206214
func uniquePaths1(grid [][]int) (ans int) {
	const mod = 1_000_000_007
	m, n := len(grid), len(grid[0])
	memo := make([][][2]int, m)
	for i := range memo {
		memo[i] = make([][2]int, n)
		for j := range memo[i] {
			memo[i][j] = [2]int{-1, -1} // -1 表示没有计算过
		}
	}
	var dfs func(int, int, int) int
	dfs = func(i, j, k int) (res int) {
		if i < 0 || j < 0 {
			return 0
		}
		if i == 0 && j == 0 {
			return 1
		}
		p := &memo[i][j][k]
		if *p != -1 { // 之前计算过
			return *p
		}
		defer func() { *p = res }() // 记忆化
		if grid[i][j] == 0 {        // 没有镜子，随便走
			return (dfs(i, j-1, 0) + dfs(i-1, j, 1)) % mod
		}
		if k == 0 { // 从下边过来
			return dfs(i-1, j, 1) // 反射到左边
		}
		// 从右边过来
		return dfs(i, j-1, 0) // 反射到上边
	}
	return dfs(m-1, n-1, 0)
}

func uniquePaths2(grid [][]int) (ans int) {
	const mod = 1_000_000_007
	m, n := len(grid), len(grid[0])
	f := make([][][2]int, m+1)
	for i := range f {
		f[i] = make([][2]int, n+1)
	}
	f[0][1] = [2]int{1, 1}
	for i, row := range grid {
		for j, x := range row {
			if x == 0 {
				f[i+1][j+1][0] = (f[i+1][j][0] + f[i][j+1][1]) % mod
				f[i+1][j+1][1] = f[i+1][j+1][0]
			} else {
				f[i+1][j+1][0] = f[i][j+1][1]
				f[i+1][j+1][1] = f[i+1][j][0]
			}
		}
	}
	return f[m][n][0]
}

func uniquePaths(grid [][]int) (ans int) {
	const mod = 1_000_000_007
	n := len(grid[0])
	f := make([][2]int, n+1)
	f[1] = [2]int{1, 1}
	for _, row := range grid {
		for j, x := range row {
			if x == 0 {
				f[j+1][0] = (f[j][0] + f[j+1][1]) % mod
				f[j+1][1] = f[j+1][0]
			} else {
				f[j+1][0] = f[j+1][1]
				f[j+1][1] = f[j][0]
			}
		}
	}
	return f[n][0]
}
