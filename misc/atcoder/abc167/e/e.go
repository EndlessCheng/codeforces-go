package main

import (
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	const mod = 998244353
	const mx int = 2e5
	F := [mx + 1]int{1}
	for i := 1; i <= mx; i++ {
		F[i] = F[i-1] * i % mod
	}
	pow := func(x, n int) int {
		res := 1
		for ; n > 0; n >>= 1 {
			if n&1 == 1 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return res
	}
	invF := [...]int{mx: pow(F[mx], mod-2)}
	for i := mx; i > 0; i-- {
		invF[i-1] = invF[i] * i % mod
	}
	c := func(n, k int) int { return F[n] * invF[k] % mod * invF[n-k] % mod }

	var n, m, k, ans int
	Fscan(_r, &n, &m, &k)
	for i := 0; i <= k; i++ {
		ans = (ans + c(n-1, i)*pow(m-1, n-1-i)) % mod
	}
	Fprint(_w, ans*m%mod)
}

func main() { run(os.Stdin, os.Stdout) }
