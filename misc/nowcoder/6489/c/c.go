package main

// github.com/EndlessCheng/codeforces-go
func getScores(n int, qs []int) []int {
	mat := make([][]int, n)
	for i := range mat {
		mat[i] = make([]int, n)
	}
	type pair struct{ x, y int }
	dir4 := [4]pair{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	var x, y, di int
	pos := make([]pair, n*n+1)
	for i := 1; i <= n*n; i++ {
		mat[x][y] = i
		pos[i] = pair{x, y}
		d := dir4[di&3]
		if xx, yy := x+d.x, y+d.y; xx < 0 || xx >= n || yy < 0 || yy >= n || mat[xx][yy] > 0 {
			di++
		}
		d = dir4[di&3]
		x += d.x
		y += d.y
	}

	_s := make([][]int, n+1)
	_s[0] = make([]int, n+1)
	for i, row := range mat {
		_s[i+1] = make([]int, n+1)
		for j, v := range row {
			_s[i+1][j+1] = _s[i+1][j] + _s[i][j+1] - _s[i][j] + v
		}
	}
	query := func(r1, c1, r2, c2 int) int {
		if r1 < 0 {
			r1 = 0
		}
		if c1 < 0 {
			c1 = 0
		}
		if r2 > n {
			r2 = n
		}
		if c2 > n {
			c2 = n
		}
		return _s[r2][c2] - _s[r2][c1] - _s[r1][c2] + _s[r1][c1]
	}

	ans := make([]int, 0, len(qs))
	sum := 0
	cur := 1
	for _, step := range qs {
		cur = (cur+step-1)%(n*n) + 1
		p := pos[cur]
		x, y := p.x, p.y
		step--
		sum1 := query(x-step, 0, x+step+1, n)
		sum2 := query(0, y-step, n, y+step+1)
		sum3 := query(x-step, y-step, x+step+1, y+step+1)
		sum = (sum + sum1 + sum2 - sum3) % (1e9 + 7)
		ans = append(ans, sum)
	}
	return ans
}
