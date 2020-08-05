package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF869C(in io.Reader, out io.Writer) {
	const m = 998244353
	const mx = 5000
	F := [mx + 1]int64{1}
	for i := 1; i <= mx; i++ {
		F[i] = F[i-1] * int64(i) % m
	}
	pow := func(x, n int64) int64 {
		res := int64(1)
		for ; n > 0; n >>= 1 {
			if n&1 == 1 {
				res = res * x % m
			}
			x = x * x % m
		}
		return res
	}
	invF := [...]int64{mx: pow(F[mx], m-2)}
	for i := mx; i > 0; i-- {
		invF[i-1] = invF[i] * int64(i) % m
	}
	C := func(n, k int) int64 { return F[n] * invF[k] % m * invF[n-k] % m }
	f := func(a, b int) (s int64) {
		for i := 0; i <= a && i <= b; i++ {
			s = (s + C(a, i)*C(b, i)%m*F[i]) % m
		}
		return
	}

	var a, b, c int
	Fscan(in, &a, &b, &c)
	Fprint(out, f(a, b)*f(b, c)%m*f(a, c)%m)
}

//func main() { CF869C(os.Stdin, os.Stdout) }
