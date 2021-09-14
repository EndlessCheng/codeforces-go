package main

// github.com/EndlessCheng/codeforces-go
func findBall(grid [][]int) []int {
	n, m := len(grid), len(grid[0])
	ans := make([]int, m)
o:
	for j := range ans {
		j0 := j
		for i, up := 0, true; i < n; up = !up {
			if up {
				dir := grid[i][j]
				if j += dir; j < 0 || j == m || grid[i][j] != dir {
					ans[j0] = -1
					continue o
				}
			} else {
				i++
			}
		}
		ans[j0] = j
	}
	return ans
}
