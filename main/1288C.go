package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1288C(_r io.Reader, _w io.Writer) {
	const p int64 = 1e9 + 7
	const mx = 1023
	F := [mx + 1]int64{1}
	for i := 1; i <= mx; i++ {
		F[i] = F[i-1] * int64(i) % p
	}
	pow := func(x, n int64) int64 {
		res := int64(1)
		for ; n > 0; n >>= 1 {
			if n&1 == 1 {
				res = res * x % p
			}
			x = x * x % p
		}
		return res
	}
	invF := [...]int64{mx: pow(F[mx], p-2)}
	for i := mx; i > 0; i-- {
		invF[i-1] = invF[i] * int64(i) % p
	}
	C := func(n, k int64) int64 { return F[n] * invF[k] % p * invF[n-k] % p }
	var n, m, ans int64
	Fscan(_r, &n, &m)
	for i := int64(1); i <= n; i++ {
		ans = (ans + C(i+m-2, m-1)*C(n-i+m, m)) % p
	}
	Fprint(_w, ans)
}

//func main() { CF1288C(os.Stdin, os.Stdout) }