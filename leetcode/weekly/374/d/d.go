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

func comb(n, k int) int {
	return fac[n] * invFac[k] % mod * invFac[n-k] % mod
}

func numberOfSequence(n int, a []int) int {
	m := len(a)
	total := n - m
	ans := comb(total, a[0]) * comb(total-a[0], n-a[m-1]-1) % mod
	total -= a[0] + n - a[m-1] - 1
	e := 0
	for i := 1; i < m; i++ {
		k := a[i] - a[i-1] - 1
		if k > 0 {
			e += k - 1
			ans = ans * comb(total, k) % mod
			total -= k
		}
	}
	return ans * pow(2, e) % mod
}

func pow(x, n int) (res int) {
	res = 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return
}
