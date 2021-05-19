package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1342E(in io.Reader, out io.Writer) {
	const mod = 998244353
	const mx int = 2e5
	F := [mx + 1]int64{1}
	for i := 1; i <= mx; i++ {
		F[i] = F[i-1] * int64(i) % mod
	}
	pow := func(x int64, n int) (res int64) {
		res = 1
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

	var n int
	var kk, ans int64
	Fscan(in, &n, &kk)
	if kk >= int64(n) {
		Fprint(out, 0)
		return
	}
	k := int(kk)

	c := n - k
	for i := c; i >= 0; i-- {
		if i&1 == c&1 {
			ans = (ans + pow(int64(i), n)*C(c, i)) % mod
		} else {
			ans = (ans + mod - pow(int64(i), n)*C(c, i)%mod) % mod
		}
	}
	ans = ans * C(n, c) % mod
	if k > 0 {
		ans = ans * 2 % mod
	}
	Fprint(out, ans)
}

//func main() { CF1342E(os.Stdin, os.Stdout) }
