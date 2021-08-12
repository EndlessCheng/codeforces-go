package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF712D(in io.Reader, out io.Writer) {
	const mod = 1e9 + 7
	var a, b, k, t int
	Fscan(in, &a, &b, &k, &t)
	t *= 2
	offset := k * t
	if b-a >= offset {
		Fprint(out, 0)
		return
	}

	n := offset*2 + k + 1
	dp := make([]int, n)
	dp[offset] = 1
	s := make([]int, n)
	for ; t > 0; t-- {
		s[0] = dp[0]
		for i := 1; i < n; i++ {
			s[i] = (s[i-1] + dp[i]) % mod
		}
		copy(dp, s[k:k*2+1])
		for i := k + 1; i <= offset*2; i++ {
			dp[i] = (s[i+k] - s[i-k-1]) % mod
		}
	}
	ans := int64(0)
	for _, v := range dp[offset-a+b+1 : offset*2+1] {
		ans += int64(v)
	}
	Fprint(out, (ans%mod+mod)%mod)
}

//func main() { CF712D(os.Stdin, os.Stdout) }
