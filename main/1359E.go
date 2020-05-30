package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1359E(in io.Reader, out io.Writer) {
	const mod int64 = 998244353
	const mx int = 5e5
	F := [mx + 1]int64{1}
	for i := 1; i <= mx; i++ {
		F[i] = F[i-1] * int64(i) % mod
	}
	pow := func(x, n int64) int64 {
		res := int64(1)
		for ; n > 0; n >>= 1 {
			if n&1 == 1 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return res
	}
	invF := [...]int64{mx: pow(F[mx], mod-2)}
	for i := mx; i > 0; i-- {
		invF[i-1] = invF[i] * int64(i) % mod
	}
	C := func(n, k int) int64 { return F[n] * invF[k] % mod * invF[n-k] % mod }

	var n, k int
	Fscan(in, &n, &k)
	ans := int64(0)
	for i := 1; n/i >= k; i++ {
		ans += C(n/i-1, k-1)
	}
	Fprint(out, ans%mod)
}

//func main() { CF1359E(os.Stdin, os.Stdout) }
