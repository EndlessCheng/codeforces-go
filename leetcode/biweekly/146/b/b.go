package main

// https://space.bilibili.com/206214
func countPathsWithXorValue(grid [][]int, k int) int {
	const mod = 1_000_000_007
	u := 0
	for _, row := range grid {
		for _, val := range row {
			u |= val
		}
	}
	if k > u {
		return 0
	}

	m, n := len(grid), len(grid[0])
	f := make([][][]int, m+1)
	for i := range f {
		f[i] = make([][]int, n+1)
		for j := range f[i] {
			f[i][j] = make([]int, u+1)
		}
	}
	f[1][1][grid[0][0]] = 1
	for i, row := range grid {
		for j, val := range row {
			for x := range u + 1 {
				f[i+1][j+1][x] += (f[i+1][j][x^val] + f[i][j+1][x^val]) % mod
			}
		}
	}
	return f[m][n][k]
}

func countPathsWithXorValue2(grid [][]int, k int) int {
	const mod = 1_000_000_007
	u := 0
	for _, row := range grid {
		for _, val := range row {
			u |= val
		}
	}
	if k > u {
		return 0
	}

	m, n := len(grid), len(grid[0])
	memo := make([][][]int, m)
	for i := range memo {
		memo[i] = make([][]int, n)
		for j := range memo[i] {
			memo[i][j] = make([]int, u+1)
			for x := range memo[i][j] {
				memo[i][j][x] = -1
			}
		}
	}
	var dfs func(int, int, int) int
	dfs = func(i, j, x int) int {
		if i < 0 || j < 0 {
			return 0
		}
		val := grid[i][j]
		if i == 0 && j == 0 {
			if x == val {
				return 1
			}
			return 0
		}
		p := &memo[i][j][x]
		if *p != -1 {
			return *p
		}
		*p = (dfs(i, j-1, x^val) + dfs(i-1, j, x^val)) % mod
		return *p
	}
	return dfs(m-1, n-1, k)
}
