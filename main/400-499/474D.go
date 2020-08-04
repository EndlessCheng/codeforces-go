package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF474D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mod int64 = 1e9 + 7
	const mx int = 1e5

	var q, k, l, r int
	Fscan(in, &q, &k)
	dp := [mx + 1]int64{}
	for i := range dp {
		if i < k {
			dp[i] = 1
		} else {
			dp[i] = (dp[i-1] + dp[i-k]) % mod
		}
	}
	for i := 1; i <= mx; i++ {
		dp[i] += dp[i-1]
	}
	for ; q > 0; q-- {
		Fscan(in, &l, &r)
		Fprintln(out, (dp[r]-dp[l-1])%mod)
	}
}

//func main() { CF474D(os.Stdin, os.Stdout) }
