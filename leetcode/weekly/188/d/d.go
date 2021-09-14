package main

func ways(mat []string, K int) (ans int) {
	n, m := len(mat), len(mat[0])
	s := make([][]int, n+1)
	s[0] = make([]int, m+1)
	for i, row := range mat {
		s[i+1] = make([]int, m+1)
		for j, v := range row {
			s[i+1][j+1] = s[i+1][j] + s[i][j+1] - s[i][j]
			if v == 'A' {
				s[i+1][j+1]++
			}
		}
	}
	q := func(r1, c1, r2, c2 int) int {
		r2++
		c2++
		return s[r2][c2] - s[r2][c1] - s[r1][c2] + s[r1][c1]
	}

	dp := make([][][]int, n)
	for i := range dp {
		dp[i] = make([][]int, m)
		for j := range dp[i] {
			dp[i][j] = make([]int, K)
			for k := range dp[i][j] {
				dp[i][j][k] = -1
			}
		}
	}
	var f func(int, int, int) int
	f = func(x int, y int, k int) (res int) {
		all := q(x, y, n-1, m-1)
		if k == 0 {
			if all > 0 {
				return 1
			}
			return 0
		}
		if all == 1 {
			return 0
		}
		dv := &dp[x][y][k]
		if *dv != -1 {
			return *dv
		}
		defer func() { *dv = res }()
		for i := x + 1; i < n; i++ {
			if c := q(x, y, i-1, m-1); 0 < c && c < all {
				res += f(i, y, k-1)
			}
		}
		for j := y + 1; j < m; j++ {
			if c := q(x, y, n-1, j-1); 0 < c && c < all {
				res += f(x, j, k-1)
			}
		}
		res %= 1e9 + 7
		return
	}
	return f(0, 0, K-1)
}
