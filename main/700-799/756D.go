package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf756D(in io.Reader, out io.Writer) {
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
	const mx = 5000
	F := [mx + 1]int{1}
	for i := 1; i <= mx; i++ {
		F[i] = F[i-1] * i % M
	}
	invF := [...]int{mx: pow(F[mx], M-2)}
	for i := mx; i > 0; i-- {
		invF[i-1] = invF[i] * i % M
	}
	C := func(n, k int) int { return F[n] * invF[k] % M * invF[n-k] % M }

	var n, ans int
	var s string
	Fscan(in, &n, &s)
	f := [26][]int{}
	for i := range f {
		f[i] = make([]int, n+1)
	}
	sumF := make([]int, n+1)
	for i, b := range s {
		b -= 'a'
		for sz := i + 1; sz > 0; sz-- {
			old := f[b][sz]
			if sz > 1 {
				f[b][sz] = (sumF[sz-1] - f[b][sz-1]) % M
			} else {
				f[b][sz] = 1
			}
			sumF[sz] = (sumF[sz] + f[b][sz] - old) % M
		}
	}

	for sz := 1; sz <= n; sz++ {
		ans = (ans + sumF[sz]*C(n-1, sz-1)) % M
	}
	Fprint(out, (ans+M)%M)
}

//func main() { cf756D(os.Stdin, os.Stdout) }
