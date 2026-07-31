package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1736E(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i])
	}

	dp := make([][][2]int, n+2)
	for i := range dp {
		dp[i] = make([][2]int, n+2)
		for j := range dp[i] {
			dp[i][j] = [2]int{-1, -1}
		}
	}

	var dfs func(int, int, int) int
	dfs = func(nw, x, y int) int {
		if nw > n {
			return 0
		}
		p := &dp[nw][x][y]
		if *p != -1 {
			return *p
		}

		res := -1
		for i := nw; i <= n; i++ {
			res = max(res, dfs(i+1, x+1, 1)+(i+1-nw)*a[nw]*y)
		}
		for i := 1; i <= x; i++ {
			ex := 0
			if i+nw <= n {
				ex = i * a[i+nw]
			}
			res = max(res, ex+max(dfs(nw+i, x-i+1, 0), dfs(nw+i, x-i, 1)))
		}

		*p = res
		return res
	}

	Fprint(out, dfs(1, 1, 1))
}

//func main() { cf1736E(bufio.NewReader(os.Stdin), os.Stdout) }
