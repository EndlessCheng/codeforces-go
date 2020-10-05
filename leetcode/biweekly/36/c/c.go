package main

// github.com/EndlessCheng/codeforces-go
func restoreMatrix(rowSum []int, colSum []int) [][]int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	n, m := len(rowSum), len(colSum)
	ans := make([][]int, n)
	for i := range ans {
		ans[i] = make([]int, m)
		for j := range ans[i] {
			ans[i][j] = min(rowSum[i], colSum[j])
			rowSum[i] -= ans[i][j]
			colSum[j] -= ans[i][j]
		}
	}
	return ans
}
