package main

import "math"

// github.com/EndlessCheng/codeforces-go
func gridGame(grid [][]int) int64 {
	ans := math.MaxInt64
	left0 := 0
	for _, v := range grid[0] {
		left0 += v
	}
	left1 := 0
	for j, v := range grid[0] { // 枚举第一个机器人在哪拐弯
		left0 -= v
		ans = min(ans, max(left0, left1)) // 第二个机器人只能取上下剩余部分中的最大值
		left1 += grid[1][j]
	}
	return int64(ans)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
