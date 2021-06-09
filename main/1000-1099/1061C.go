package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1061C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	const mod int = 1e9 + 7
	const mx int = 1e6

	var n, v, ans int
	Fscan(in, &n)
	ds := [mx + 1][]int{}
	for i := 1; i <= n; i++ {
		for j := i; j <= mx; j += i {
			ds[j] = append(ds[j], i)
		}
	}
	dp := make([]int, n+1)
	dp[0] = 1
	for ; n > 0; n-- {
		Fscan(in, &v)
		for j := len(ds[v]) - 1; j >= 0; j-- {
			d := ds[v][j]
			ans = (ans + dp[d-1]) % mod
			dp[d] = (dp[d] + dp[d-1]) % mod
		}
	}
	Fprint(out, ans)
}

//func main() { CF1061C(os.Stdin, os.Stdout) }
