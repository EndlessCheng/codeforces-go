package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1906H(in io.Reader, out io.Writer) {
	const mod = 998244353
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
	const mx int = 2e5
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

	var n, m int
	var s, t string
	Fscan(in, &n, &m, &s, &t)
	var cntS, cntT [26]int
	for _, b := range s {
		cntS[b-'A']++
	}
	for _, b := range t {
		cntT[b-'A']++
	}
	if cntS[25] > cntT[25] {
		Fprint(out, 0)
		return
	}
	cntT[25] -= cntS[25]

	f := make([]int, n+1)
	f[0] = 1
	sum := make([]int, n+2)
	for i, cs := range cntS[:25] {
		for j, v := range f {
			sum[j+1] = sum[j] + v
		}
		clear(f)
		for j := max(cs-cntT[i], 0); j <= min(cs, cntT[i+1]); j++ {
			f[j] = sum[min(cntT[i]-cs+j, n)+1] % mod * C(cs, j) % mod
		}
	}

	sumF := 0
	for _, v := range f {
		sumF += v
	}
	perm := F[n]
	for _, c := range cntS {
		perm = perm * invF[c] % mod
	}
	Fprint(out, sumF%mod*perm%mod)
}

//func main() { cf1906H(bufio.NewReader(os.Stdin), os.Stdout) }
