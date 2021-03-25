package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF839D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	const mod int64 = 1e9 + 7
	const mx int = 1e6
	p2 := [mx + 1]int64{1}
	for i := 1; i <= mx; i++ {
		p2[i] = p2[i-1] * 2 % mod
	}

	cnt := [mx + 1]int{}
	var n, v int
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &v)
		cnt[v]++
	}
	ans := int64(0)
	dp := [mx + 1]int64{}
	for i := mx; i > 1; i-- {
		c := 0
		for j := i; j <= mx; j += i {
			c += cnt[j]
		}
		if c > 0 {
			dp[i] = int64(c) * p2[c-1]
			for j := 2 * i; j <= mx; j += i {
				dp[i] -= dp[j]
			}
			dp[i] = (dp[i]%mod + mod) % mod
			ans += dp[i] * int64(i) % mod
		}
	}
	Fprint(out, ans%mod)
}

//func main() { CF839D(os.Stdin, os.Stdout) }
