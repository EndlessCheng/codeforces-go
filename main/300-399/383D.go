package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF383D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	const mod int = 1e9 + 7
	var n, offset, s, ans int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		offset += a[i]
	}
	dp := make([]int, offset*2+1)
	dp[offset] = 1
	for _, v := range a {
		ans = (ans + (dp[offset+v]+dp[offset-v])%mod) % mod
		ndp := make([]int, len(dp))
		for i := offset - s; i <= offset+s; i++ {
			ndp[i-v] = (ndp[i-v] + dp[i]) % mod
			ndp[i+v] = (ndp[i+v] + dp[i]) % mod
		}
		ndp[offset] = (ndp[offset] + 1) % mod
		dp = ndp
		s += v
	}
	Fprint(out, ans%mod)
}

//func main() { CF383D(os.Stdin, os.Stdout) }
