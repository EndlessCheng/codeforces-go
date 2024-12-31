package main

// https://space.bilibili.com/206214
const mod = 1_000_000_007
const mx = 100_000

var F, invF [mx]int

func init() {
	F[0] = 1
	for i := 1; i < mx; i++ {
		F[i] = F[i-1] * i % mod
	}
	invF[mx-1] = pow(F[mx-1], mod-2)
	for i := mx - 1; i > 0; i-- {
		invF[i-1] = invF[i] * i % mod
	}
}

func pow(x, n int) int {
	res := 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return res
}

func comb(n, m int) int {
	return F[n] * invF[m] % mod * invF[n-m] % mod
}

func countGoodArrays(n, m, k int) int {
	return comb(n-1, k) * m % mod * pow(m-1, n-k-1) % mod
}
