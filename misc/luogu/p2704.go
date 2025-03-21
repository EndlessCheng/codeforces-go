package main

import (
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func p2704(in io.Reader, out io.Writer) {
	var n, m int
	Fscan(in, &n, &m)
	a := make([]string, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	pow3 := make([]int, m)
	pow3[0] = 1
	for i := 1; i < m; i++ {
		pow3[i] = pow3[i-1] * 3
	}

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, pow3[m-1]*3)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	var dfs func(int, int) int
	dfs = func(i, s int) (res int) {
		if i < 0 {
			return
		}
		dv := &dp[i][s]
		if *dv >= 0 {
			return *dv
		}
		var f func(int, int, int)
		f = func(j, t, c int) {
			if j < 0 {
				res = max(res, dfs(i-1, t)+c)
				return
			}
			if v := s / pow3[j] % 3; v > 0 {
				f(j-1, t*3+v-1, c)
				return
			}
			f(j-1, t*3, c)
			if t%3 < 2 && t/3%3 < 2 && a[i][j] == 'P' {
				f(j-1, t*3+2, c+1)
			}
		}
		f(m-1, 0, 0)
		*dv = res
		return
	}
	Fprint(out, dfs(n-1, 0))
}

//func main() { p2704(bufio.NewReader(os.Stdin), os.Stdout) }
