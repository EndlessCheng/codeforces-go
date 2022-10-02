package main

// https://space.bilibili.com/206214
func maxSum(grid [][]int) (ans int) {
	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[i])-1; j++ {
			ans = max(ans, grid[i-1][j-1]+grid[i-1][j]+grid[i-1][j+1]+grid[i][j]+grid[i+1][j-1]+grid[i+1][j]+grid[i+1][j+1])
		}
	}
	return
}

func max(a, b int) int { if b > a { return b }; return a }
