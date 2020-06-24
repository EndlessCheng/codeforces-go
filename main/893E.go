package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF893E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mx = 1 << 20
	lpf := [mx + 1]int{1: 1}
	for i := 2; i <= mx; i++ {
		if lpf[i] == 0 {
			for j := i; j <= mx; j += i {
				if lpf[j] == 0 {
					lpf[j] = i
				}
			}
		}
	}
	const mod int64 = 1e9 + 7
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
	pow2 := [mx + 1]int64{1}
	for i := 1; i <= mx; i++ {
		pow2[i] = pow2[i-1] << 1 % mod
	}

	var t, n, m int
	for Fscan(in, &t); t > 0; t-- {
		ans := int64(1)
		Fscan(in, &n, &m)
		for n > 1 {
			p, e := lpf[n], 1
			for n /= p; lpf[n] == p; n /= p {
				e++
			}
			ans = ans * C(e+m-1, e) % mod
		}
		Fprintln(out, ans*pow2[m-1]%mod)
	}
}

//func main() { CF893E(os.Stdin, os.Stdout) }
