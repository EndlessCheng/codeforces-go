package main

// github.com/EndlessCheng/codeforces-go
func canMouseWin(a []string, catJump, mouseJump int) bool {
	var mx, my, cx, cy, fx, fy int
	for i, r := range a {
		for j, b := range r {
			if b == 'M' {
				mx, my = i, j
			} else if b == 'C' {
				cx, cy = i, j
			} else if b == 'F' {
				fx, fy = i, j
			}
		}
	}

	type pair struct{ x, y int }
	dir4 := []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	n, m := len(a), len(a[0])
	dp := make([][][2]int, n*m*n*m)
	for i := range dp {
		dp[i] = make([][2]int, n*m+10)
		for j := range dp[i] {
			dp[i][j] = [2]int{-1, -1}
		}
	}
	var f func(mx, my, cx, cy, who, dep int) int
	f = func(mx, my, cx, cy, who, dep int) (res int) {
		if dep >= n*m+10 || cx == mx && cy == my || cx == fx && cy == fy {
			return who
		}
		if mx == fx && my == fy {
			return who ^ 1
		}
		dv := &dp[(mx*m+my)*n*m+cx*m+cy][dep][who]
		if *dv != -1 {
			return *dv
		}
		defer func() { *dv = res }()
		if who == 0 {
			for _, d := range dir4 {
				for j := 0; j <= mouseJump; j++ {
					if x, y := mx+d.x*j, my+d.y*j; 0 <= x && x < n && 0 <= y && y < m && a[x][y] != '#' {
						if f(x, y, cx, cy, 1, dep+1) == 0 {
							return 1
						}
					} else {
						break
					}
				}
			}
		} else {
			for _, d := range dir4 {
				for j := 0; j <= catJump; j++ {
					if x, y := cx+d.x*j, cy+d.y*j; 0 <= x && x < n && 0 <= y && y < m && a[x][y] != '#' {
						if f(mx, my, x, y, 0, dep+1) == 0 {
							return 1
						}
					} else {
						break
					}
				}
			}
		}
		return 0
	}
	return f(mx, my, cx, cy, 0, 0) == 1
}
