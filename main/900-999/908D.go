package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF908D(in io.Reader, out io.Writer) {
	const mod int64 = 1e9 + 7
	pow := func(x, n int64) int64 {
		res := int64(1)
		for ; n > 0; n >>= 1 {
			if n&1 > 0 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return res
	}

	var k, a, b int64
	Fscan(in, &k, &a, &b)
	pa := a * pow(a+b, mod-2) % mod
	pb := b * pow(a+b, mod-2) % mod
	ab := a * pow(b, mod-2) % mod
	dp := make([]int64, k)
	for i := k; i > 0; i-- {
		for j := k - 1; j >= 0; j-- {
			if i+j >= k {
				dp[j] = (i + j + ab) % mod
			} else {
				dp[j] = (pa*dp[j] + pb*dp[i+j]) % mod
			}
		}
	}
	Fprint(out, dp[0])
}

//func main() { CF908D(os.Stdin, os.Stdout) }
