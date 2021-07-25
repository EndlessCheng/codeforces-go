package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1154F(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	var n, m, k, x, y int
	Fscan(in, &n, &m, &k)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	sort.Ints(a)
	sum := make([]int, k+1)
	for i, v := range a[:k] {
		sum[i+1] = sum[i] + v
	}
	free := make([]int, k+1)
	for ; m > 0; m-- {
		if Fscan(in, &x, &y); x <= k && y > free[x] {
			free[x] = y
		}
	}
	dp := make([]int, k+1)
	for i := 1; i <= k; i++ {
		dp[i] = 1e9
		for j := 0; j < i; j++ {
			dp[i] = min(dp[i], dp[j]+sum[i]-sum[j+free[i-j]])
		}
	}
	Fprint(out, dp[k])
}

//func main() { CF1154F(os.Stdin, os.Stdout) }
