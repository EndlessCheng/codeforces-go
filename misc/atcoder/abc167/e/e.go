package main

import (
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	const mod = 998244353
	const mx = 2e5
	F := [mx + 1]int{1}
	for i := 1; i <= mx; i++ {
		F[i] = F[i-1] * i % mod
	}
	pow := func(x, n int) int {
		r := 1
		for ; n > 0; n >>= 1 {
			if n&1 == 1 {
				r = r * x % mod
			}
			x = x * x % mod
		}
		return r
	}
	c := func(n, k int) int { return F[n] * pow(F[k]*F[n-k]%mod, mod-2) % mod }

	var n, m, k, ans int
	Fscan(_r, &n, &m, &k)
	for i := 0; i <= k; i++ {
		ans = (ans + c(n-1, i)*pow(m-1, n-1-i)) % mod
	}
	Fprint(_w, ans*m%mod)
}

func main() { run(os.Stdin, os.Stdout) }
