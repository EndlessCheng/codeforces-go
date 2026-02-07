package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1096E(in io.Reader, out io.Writer) {
	const mod = 998244353
	const mx = 5100
	var inv, F, invF [mx]int
	inv[1] = 1
	F[0], F[1] = 1, 1
	invF[0], invF[1] = 1, 1
	for i := 2; i < mx; i++ {
		inv[i] = (mod - mod/i) * inv[mod%i] % mod
		F[i] = F[i-1] * i % mod
		invF[i] = invF[i-1] * inv[i] % mod
	}

	var p, s, r, ans int
	Fscan(in, &p, &s, &r)
	for i := 1; i <= p && s-r*i >= 0; i++ {
		ans += (i%2*2 - 1) * F[s-r*i+p-1] * invF[s-r*i] % mod * invF[i] % mod * invF[p-i] % mod
	}
	ans = (ans%mod + mod) % mod
	ans = ans * F[p-1] % mod * F[s-r] % mod * invF[s-r+p-1] % mod
	Fprint(out, ans)
}

//func main() { cf1096E(os.Stdin, os.Stdout) }
