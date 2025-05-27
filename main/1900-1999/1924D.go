package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1924D(in io.Reader, out io.Writer) {
	const M = 1_000_000_007
	pow := func(x, n int) int {
		res := 1
		for ; n > 0; n /= 2 {
			if n%2 > 0 {
				res = res * x % M
			}
			x = x * x % M
		}
		return res
	}
	const mx = 4000
	F := [mx + 1]int{1}
	for i := 1; i <= mx; i++ {
		F[i] = F[i-1] * i % M
	}
	invF := [...]int{mx: pow(F[mx], M-2)}
	for i := mx; i > 0; i-- {
		invF[i-1] = invF[i] * i % M
	}
	C := func(n, k int) int { return F[n] * invF[k] % M * invF[n-k] % M }

	var T, n, m, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m, &k)
		if k > min(n, m) {
			Fprintln(out, 0)
		} else {
			Fprintln(out, (C(n+m, k)-C(n+m, k-1)+M)%M)
		}
	}
}

//func main() { cf1924D(os.Stdin, os.Stdout) }
