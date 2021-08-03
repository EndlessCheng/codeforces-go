package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1140E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	const mod = 998244353
	pow := func(x int64, n int) (res int64) {
		res = 1
		for ; n > 0; n >>= 1 {
			if n&1 > 0 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return
	}

	var n, v int
	var k int64
	Fscan(in, &n, &k)
	dp := make([][2]int64, n)
	dp[0][0] = 1
	for i := 1; i < n; i++ {
		dp[i][0] = (dp[i-1][0]*(k-2) + dp[i-1][1]) % mod
		dp[i][1] = dp[i-1][0] * (k - 1) % mod
	}

	a := [2][]int{}
	for ; n > 0; n-- {
		Fscan(in, &v)
		a[n&1] = append(a[n&1], v)
	}

	f := func(a []int) int64 {
		ans := int64(1)
		for i, n := 0, len(a); i < n; i++ {
			if a[i] > 0 {
				if i > 0 && a[i] == a[i-1] {
					return 0
				}
				continue
			}
			st := i
			for ; i < n && a[i] < 0; i++ {
			}
			if st == 0 {
				if i == n {
					ans = ans * pow(k-1, n-1) % mod * k % mod
				} else {
					ans = ans * pow(k-1, i) % mod
				}
			} else if i == n {
				ans = ans * pow(k-1, n-st) % mod
			} else if a[st-1] != a[i] {
				ans = ans * dp[i-st][0] % mod
			} else {
				ans = ans * dp[i-st][1] % mod
			}
		}
		return ans
	}
	Fprint(out, f(a[0])*f(a[1])%mod)
}

//func main() { CF1140E(os.Stdin, os.Stdout) }
