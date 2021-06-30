package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF453B(in io.Reader, out io.Writer) {
	primes := [...]int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53}
	const mx = len(primes)
	const mx2 = 1 << mx
	masks := [59]int{}
	for i := 2; i < 59; i++ {
		for j, p := range primes {
			if i%p == 0 {
				masks[i] |= 1 << j
			}
		}
	}
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	var n, v int
	Fscan(in, &n)
	dp := make([][mx2]int, n+1)
	for i := range dp {
		for j := 0; j < mx2; j++ {
			dp[i][j] = 1e9
		}
	}
	dp[0][0] = 0
	type pair struct{ j, v int }
	fa := make([][mx2]pair, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &v)
		for j := 1; j == 1 || j < v*2-1; j++ {
			m := masks[j]
			c := m ^ (mx2 - 1)
			for sub, ok := c, true; ok; ok = sub != c {
				if s := dp[i-1][sub] + abs(j-v); s < dp[i][m|sub] {
					dp[i][m|sub] = s
					fa[i][m|sub] = pair{sub, j}
				}
				sub = (sub - 1) & c
			}
		}
	}
	miJ := 0
	for j, v := range dp[n] {
		if v < dp[n][miJ] {
			miJ = j
		}
	}
	ans := make([]int, n)
	for i, j := n, miJ; i > 0; i-- {
		ans[i-1] = fa[i][j].v
		j = fa[i][j].j
	}
	for _, v := range ans {
		Fprint(out, v, " ")
	}
}

//func main() { CF453B(os.Stdin, os.Stdout) }
