package main

// https://space.bilibili.com/206214
const mod = 1_000_000_007
const mx = 100_000

var fac, invFac [mx]int

func init() {
	fac[0] = 1
	for i := 1; i < mx; i++ {
		fac[i] = fac[i-1] * i % mod
	}
	invFac[mx-1] = pow(fac[mx-1], mod-2)
	for i := mx - 1; i > 0; i-- {
		invFac[i-1] = invFac[i] * i % mod
	}
}

func numberOfSequence(n int, a []int) int {
	m := len(a)
	ans := fac[n-m] * invFac[a[0]] % mod * invFac[n-1-a[m-1]] % mod
	e := 0
	for i := 1; i < m; i++ {
		k := a[i] - a[i-1] - 1
		if k > 0 {
			e += k - 1
			ans = ans * invFac[k] % mod
		}
	}
	return ans * pow(2, e) % mod
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
