package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF300C(in io.Reader, out io.Writer) {
	const m int64 = 1e9 + 7
	const mx int = 1e6
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
	var a, b, n int
	Fscan(in, &a, &b, &n)
	C := func(k int) int64 { return F[n] * invF[k] % m * invF[n-k] % m }
	ans := int64(0)
	d := n * a
	for i := 0; i <= n; i++ {
		v := d
		for ; v > 0; v /= 10 {
			if v%10 != a && v%10 != b {
				break
			}
		}
		if v == 0 {
			ans = (ans + C(i)) % m
		}
		d += b - a
	}
	Fprint(out, ans)
}

//func main() { CF300C(os.Stdin, os.Stdout) }
