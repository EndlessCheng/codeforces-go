package main

import "math/bits"

const mod, mx int = 1e9 + 7, 900
var F, invF [mx + 1]int
func init() {
	F[0] = 1
	for i := 1; i <= mx; i++ {
		F[i] = F[i-1] * i % mod
	}
	invF[mx] = pow(F[mx], mod-2)
	for i := mx; i > 0; i-- {
		invF[i-1] = invF[i] * i % mod
	}
}
func pow(x, n int) (res int) {
	res = 1
	for ; n > 0; n >>= 1 {
		if n&1 == 1 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return res
}
func C(n, k int) int { return F[n] * invF[k] % mod * invF[n-k] % mod }

// github.com/EndlessCheng/codeforces-go
func solve(N, M, k int) (ans int) {
	for s := uint(0); s < 1<<4; s++ {
		n, m := N, M
		for i := 0; i < 4; i++ {
			if s>>i&1 > 0 {
				if i&1 > 0 {
					n--
				} else {
					m--
				}
			}
		}
		if n*m < k {
			continue
		}
		if bits.OnesCount(s)&1 == 0 {
			ans += C(n*m, k)
		} else {
			ans -= C(n*m, k)
		}
	}
	return (ans%mod + mod) % mod
}
