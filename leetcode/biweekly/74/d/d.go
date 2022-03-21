package main

// github.com/EndlessCheng/codeforces-go
func minimumWhiteTiles(floor string, n, carpetLen int) int {
	m := len(floor)
	f := make([][]int, n+1)
	f[0] = make([]int, m)
	f[0][0] = int(floor[0] & 1)
	for i := 1; i < m; i++ {
		f[0][i] = f[0][i-1] + int(floor[i]&1)
	}
	for i := 1; i <= n; i++ {
		f[i] = make([]int, m)
		// j < carpetLen 的 f[i][j] 均为 0
		for j := carpetLen; j < m; j++ {
			f[i][j] = min(f[i][j-1]+int(floor[j]&1), f[i-1][j-carpetLen])
		}
	}
	return f[n][m-1]
}

func min(a, b int) int { if a > b { return b }; return a }
