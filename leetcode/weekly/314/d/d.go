package main

// https://space.bilibili.com/206214
func numberOfPaths1(grid [][]int, k int) int {
	const mod = 1_000_000_007
	m, n := len(grid), len(grid[0])
	memo := make([][][]int, m)
	for i := range memo {
		memo[i] = make([][]int, n)
		for j := range memo[i] {
			memo[i][j] = make([]int, k)
			for s := range memo[i][j] {
				memo[i][j][s] = -1
			}
		}
	}

	var dfs func(int, int, int) int
	dfs = func(i, j, s int) int {
		if i < 0 && j == 0 {
			if s == 0 {
				return 1
			}
			return 0
		}
		if i < 0 || j < 0 {
			return 0
		}
		newS := (s + grid[i][j]) % k
		p := &memo[i][j][s]
		if *p == -1 {
			*p = (dfs(i-1, j, newS) + dfs(i, j-1, newS)) % mod
		}
		return *p
	}

	return dfs(m-1, n-1, 0)
}

func numberOfPaths2(grid [][]int, k int) int {
	const mod = 1_000_000_007
	m, n := len(grid), len(grid[0])
	f := make([][][]int, m+1)
	for i := range f {
		f[i] = make([][]int, n+1)
		for j := range f[i] {
			f[i][j] = make([]int, k)
		}
	}
	f[0][1][0] = 1
	for i, row := range grid {
		for j, x := range row {
			for s := range k {
				newS := (s + x) % k
				f[i+1][j+1][s] = (f[i][j+1][newS] + f[i+1][j][newS]) % mod
			}
		}
	}
	return f[m][n][0]
}

func numberOfPaths(grid [][]int, k int) int {
	const mod = 1_000_000_007
	n := len(grid[0])
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, k)
	}
	f[1][0] = 1
	newF := make([]int, k) // 避免在循环内反复创建 []int
	for _, row := range grid {
		for j, x := range row {
			for s := range k {
				newS := (s + x) % k
				newF[s] = (f[j+1][newS] + f[j][newS]) % mod
			}
			copy(f[j+1], newF) // 复制到 f[j+1] 中
		}
	}
	return f[n][0]
}
