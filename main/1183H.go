package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1183H(_r io.Reader, _w io.Writer) {
	var n int
	var k, ans int64
	var s []byte
	Fscan(bufio.NewReader(_r), &n, &k, &s)
	prev := [26]int{}
	dp := make([][]int64, n+1)
	for i := range dp {
		dp[i] = make([]int64, n+1)
	}
	dp[0][0] = 1
	for i := 1; i <= n; i++ {
		c := s[i-1] - 'a'
		dp[i][0] = 1
		for j := 1; j <= i; j++ {
			dp[i][j] = dp[i-1][j] + dp[i-1][j-1]
			if p := prev[c]; p > 0 {
				dp[i][j] -= dp[p-1][j-1]
			}
		}
		prev[c] = i
	}
	k--
	for l := n - 1; l >= 0; l-- {
		cnt := dp[n][l]
		if cnt > k {
			ans += int64(n-l) * k
			k = 0
			break
		}
		k -= cnt
		ans += int64(n-l) * cnt
	}
	if k > 0 {
		ans = -1
	}
	Fprint(_w, ans)
}

//func main() { CF1183H(os.Stdin, os.Stdout) }
