package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1369D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mod int = 1e9 + 7
	const mx int = 2e6
	dp := [mx + 1]int{}
	for i := 3; i <= mx; i++ {
		dp[i] = dp[i-1] + 2*dp[i-2]%mod
		if i%3 == 0 {
			dp[i] += 4
		}
		dp[i] %= mod
	}

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		Fprintln(out, dp[n])
	}
}

//func main() { CF1369D(os.Stdin, os.Stdout) }
