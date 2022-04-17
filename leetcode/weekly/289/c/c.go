package main

// github.com/EndlessCheng/codeforces-go
var c25 [1001][2]int

func init() {
	// 预处理：递推算出每个数的因子 2 的个数和因子 5 的个数
	for i := 2; i <= 1000; i++ {
		if i%2 == 0 { c25[i][0] = c25[i/2][0] + 1 }
		if i%5 == 0 { c25[i][1] = c25[i/5][1] + 1 }
	}
}

func maxTrailingZeros(grid [][]int) (ans int) {
	m, n := len(grid), len(grid[0])
	s := make([][][2]int, m)
	for i, row := range grid {
		s[i] = make([][2]int, n+1)
		for j, v := range row {
			s[i][j+1][0] = s[i][j][0] + c25[v][0] // 每行的因子 2 的前缀和
			s[i][j+1][1] = s[i][j][1] + c25[v][1] // 每行的因子 5 的前缀和
		}
	}

	for j := 0; j < n; j++ {
		s2, s5 := 0, 0
		for i, row := range grid { // 从上往下，枚举左拐还是右拐
			s2 += c25[row[j]][0]
			s5 += c25[row[j]][1]
			ans = max(ans, max(min(s2+s[i][j][0], s5+s[i][j][1]), min(s2+s[i][n][0]-s[i][j+1][0], s5+s[i][n][1]-s[i][j+1][1])))
		}
		s2, s5 = 0, 0
		for i := m - 1; i >= 0; i-- { // 从下往上，枚举左拐还是右拐
			s2 += c25[grid[i][j]][0]
			s5 += c25[grid[i][j]][1]
			ans = max(ans, max(min(s2+s[i][j][0], s5+s[i][j][1]), min(s2+s[i][n][0]-s[i][j+1][0], s5+s[i][n][1]-s[i][j+1][1])))
		}
	}
	return
}

func max(a, b int) int { if a < b { return b }; return a }
func min(a, b int) int { if a > b { return b }; return a }
