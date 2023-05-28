package main

// https://space.bilibili.com/206214
func differenceOfDistinctValues(grid [][]int) [][]int {
	m, n := len(grid), len(grid[0])
	ans := make([][]int, m)
	for i := range ans {
		ans[i] = make([]int, n)
	}
	for s := 1; s < m+n; s++ {
		minJ := max(0, n-s)
		maxJ := min(n-1, n-s+m-1)
		// topLeft
		set := map[int]struct{}{}
		for j := minJ; j < maxJ; j++ {
			i := s + j - n
			set[grid[i][j]] = struct{}{}
			ans[i+1][j+1] = len(set)
		}
		// bottomRight
		set = map[int]struct{}{}
		for j := maxJ; j > minJ; j-- {
			i := s + j - n
			set[grid[i][j]] = struct{}{}
			ans[i-1][j-1] = abs(ans[i-1][j-1] - len(set))
		}
	}
	return ans
}

func abs(x int) int { if x < 0 { return -x }; return x }
func min(a, b int) int { if b < a { return b }; return a }
func max(a, b int) int { if b > a { return b }; return a }
