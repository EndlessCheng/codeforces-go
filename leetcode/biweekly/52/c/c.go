package main

// github.com/EndlessCheng/codeforces-go
func rotateTheBox(a [][]byte) [][]byte {
	n, m := len(a), len(a[0])
	ans := make([][]byte, m)
	for i := range ans {
		ans[i] = make([]byte, n)
		for j := range ans[i] {
			ans[i][j] = '.'
		}
	}
	for i, r := range a {
		for j := 0; j < m; j++ {
			c := 0
			for ; j < m && r[j] != '*'; j++ {
				if r[j] == '#' {
					c++
				}
			}
			if j < m {
				ans[j][n-1-i] = '*'
			}
			for k := j - 1; c > 0; k-- {
				ans[k][n-1-i] = '#'
				c--
			}
		}
	}
	return ans
}
