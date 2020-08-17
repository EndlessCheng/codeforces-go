package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1312D(_r io.Reader, _w io.Writer) {
	const mx int = 2e5
	const mod int64 = 998244353
	fact := [mx + 1]int64{1}
	for i := 1; i <= mx; i++ {
		fact[i] = fact[i-1] * int64(i) % mod
	}
	pow := func(x, n int64) int64 {
		x %= mod
		res := int64(1)
		for ; n > 0; n >>= 1 {
			if n&1 == 1 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return res
	}
	inv := func(a int64) int64 { return pow(a, mod-2) }
	div := func(a, b int64) int64 { return a * inv(b) % mod }
	comb := func(n, k int64) int64 { return div(fact[n], fact[k]*fact[n-k]%mod) }

	var n, m int64
	Fscan(_r, &n, &m)
	if n == 2 {
		Fprint(_w, 0)
	} else {
		Fprint(_w, comb(m, n-1)*(n-2)%mod*pow(2, n-3)%mod)
	}
}

//func main() { CF1312D(os.Stdin, os.Stdout) }
