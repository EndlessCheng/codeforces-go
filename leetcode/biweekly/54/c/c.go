package main

// github.com/EndlessCheng/codeforces-go
func largestMagicSquare(a [][]int) int {
	n, m := len(a), len(a[0])
	rs := make([][]int, n)
	cs := make([][]int, n+1)
	ds := make([][]int, n+1)
	as := make([][]int, n+1)
	for i := range cs {
		cs[i] = make([]int, m)
		ds[i] = make([]int, m+1)
		as[i] = make([]int, m+1)
	}
	for i, row := range a {
		rs[i] = make([]int, m+1)
		for j, v := range row {
			rs[i][j+1] = rs[i][j] + v
			cs[i+1][j] = cs[i][j] + v
			ds[i+1][j+1] = ds[i][j] + v
			as[i+1][j] = as[i][j+1] + v
		}
	}

	for k := min(n, m); ; k-- {
		for r := k; r <= n; r++ {
		o:
			for c := k; c <= m; c++ {
				s := ds[r][c] - ds[r-k][c-k]
				if as[r][c-k]-as[r-k][c] != s {
					continue
				}
				for _, row := range rs[r-k : r] {
					if row[c]-row[c-k] != s {
						continue o
					}
				}
				for j := c - k; j < c; j++ {
					if cs[r][j]-cs[r-k][j] != s {
						continue o
					}
				}
				return k
			}
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
