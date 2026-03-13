package main

import "math"

// github.com/EndlessCheng/codeforces-go
func maxProductPath1(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	memo := make([][][2]int, m)
	for i := range memo {
		memo[i] = make([][2]int, n)
		for j := range memo[i] {
			memo[i][j] = [2]int{math.MinInt, math.MinInt}
		}
	}

	var dfs func(int, int) (int, int)
	dfs = func(i, j int) (int, int) {
		x := grid[i][j]
		if i == 0 && j == 0 {
			return x, x
		}

		p := &memo[i][j]
		if p[0] != math.MinInt {
			return p[0], p[1]
		}

		resMin := math.MaxInt
		resMax := math.MinInt
		if i > 0 {
			mn, mx := dfs(i-1, j)
			resMin = min(mn*x, mx*x)
			resMax = max(mn*x, mx*x)
		}
		if j > 0 {
			mn, mx := dfs(i, j-1)
			resMin = min(resMin, mn*x, mx*x)
			resMax = max(resMax, mn*x, mx*x)
		}

		p[0], p[1] = resMin, resMax
		return resMin, resMax
	}

	_, ans := dfs(m-1, n-1)
	if ans < 0 {
		return -1
	}
	return ans % 1_000_000_007
}

func maxProductPath2(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	f := make([][][2]int, m)
	for i := range f {
		f[i] = make([][2]int, n)
	}

	for i, row := range grid {
		for j, x := range row {
			if i == 0 && j == 0 {
				f[0][0] = [2]int{x, x}
				continue
			}

			resMin := math.MaxInt
			resMax := math.MinInt
			if i > 0 {
				mn, mx := f[i-1][j][0], f[i-1][j][1]
				resMin = min(mn*x, mx*x)
				resMax = max(mn*x, mx*x)
			}
			if j > 0 {
				mn, mx := f[i][j-1][0], f[i][j-1][1]
				resMin = min(resMin, mn*x, mx*x)
				resMax = max(resMax, mn*x, mx*x)
			}

			f[i][j] = [2]int{resMin, resMax}
		}
	}

	ans := f[m-1][n-1][1]
	if ans < 0 {
		return -1
	}
	return ans % 1_000_000_007
}

func maxProductPath(grid [][]int) int {
	n := len(grid[0])
	fMin := make([]int, n)
	fMax := make([]int, n)

	for i, row := range grid {
		for j, x := range row {
			if i == 0 && j == 0 {
				fMin[0], fMax[0] = x, x
				continue
			}

			resMin := math.MaxInt
			resMax := math.MinInt
			if i > 0 {
				mn, mx := fMin[j], fMax[j]
				resMin = min(mn*x, mx*x)
				resMax = max(mn*x, mx*x)
			}
			if j > 0 {
				mn, mx := fMin[j-1], fMax[j-1]
				resMin = min(resMin, mn*x, mx*x)
				resMax = max(resMax, mn*x, mx*x)
			}

			fMin[j] = resMin
			fMax[j] = resMax
		}
	}

	ans := fMax[n-1]
	if ans < 0 {
		return -1
	}
	return ans % 1_000_000_007
}
