package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf708E(in io.Reader, out io.Writer) {
	const mod = 1_000_000_007
	pow := func(x, n int) int {
		res := 1
		for ; n > 0; n /= 2 {
			if n%2 > 0 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return res
	}
	const mx int = 1e5
	F := [mx + 1]int{1}
	for i := 1; i <= mx; i++ {
		F[i] = F[i-1] * i % mod
	}
	invF := [...]int{mx: pow(F[mx], mod-2)}
	for i := mx; i > 0; i-- {
		invF[i-1] = invF[i] * i % mod
	}
	C := func(n, k int) int {
		return F[n] * invF[k] % mod * invF[n-k] % mod
	}

	var n, m, a, b, k int
	Fscan(in, &n, &m, &a, &b, &k)
	p0 := a * pow(b, mod-2) % mod
	p := make([]int, m+1)
	for i := range min(m, k) + 1 {
		p[i] = C(k, i) * pow(p0, i) % mod * pow(1-p0, k-i) % mod
	}
	sumP := make([]int, m+1)
	sumP[0] = p[0]
	for i := 1; i <= m; i++ {
		sumP[i] = (sumP[i-1] + p[i]) % mod
	}

	f := make([]int, m+1)
	f[m] = 1
	sumF := make([]int, m+1)
	for range n {
		for j := 1; j <= m; j++ {
			sumF[j] = (sumF[j-1] + f[j]) % mod
		}
		sumFP := 0
		for j := 1; j <= m; j++ {
			f[j] = ((sumF[m]-sumF[m-j])*sumP[j-1] - sumFP) % mod * p[m-j] % mod
			sumFP = (sumFP + sumF[j]*p[j]) % mod
		}
	}

	ans := 0
	for _, v := range f {
		ans += v
	}
	Fprint(out, (ans%mod+mod)%mod)
}

//func main() { cf708E(os.Stdin, os.Stdout) }
