package main

// https://space.bilibili.com/206214
const mod = 1_000_000_007
const mx = 2000
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

func valueAfterKSeconds(n, k int) int {
	return F[n+k-1] * invF[n-1] % mod * invF[k] % mod
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
