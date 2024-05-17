package main

// github.com/EndlessCheng/codeforces-go
func matrixBlockSum(a [][]int, k int) (ans [][]int) {
	n, m := len(a), len(a[0])
	s := make([][]int, n+1)
	s[0] = make([]int, m+1)
	for i, row := range a {
		s[i+1] = make([]int, m+1)
		for j, v := range row {
			s[i+1][j+1] = s[i+1][j] + s[i][j+1] - s[i][j] + v
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
			ans[i][j] = s[r2][c2] - s[r2][c1] - s[r1][c2] + s[r1][c1]
		}
	}
	return
}
