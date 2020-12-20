package main

// github.com/EndlessCheng/codeforces-go
func matrixBlockSum(a [][]int, k int) (ans [][]int) {
	n, m := len(a), len(a[0])
	sum2d := make([][]int, n+1)
	sum2d[0] = make([]int, m+1)
	for i, row := range a {
		sum2d[i+1] = make([]int, m+1)
		for j, v := range row {
			sum2d[i+1][j+1] = sum2d[i+1][j] + sum2d[i][j+1] - sum2d[i][j] + v
		}
	}
	ans = make([][]int, n)
	for i := range ans {
		ans[i] = make([]int, m)
		for j := range ans[i] {
			r1 := max(0, i-k)
			c1 := max(0, j-k)
			r2 := min(n, i+k+1)
			c2 := min(m, j+k+1)
			ans[i][j] = sum2d[r2][c2] - sum2d[r2][c1] - sum2d[r1][c2] + sum2d[r1][c1]
		}
	}
	return
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
