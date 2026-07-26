package main

// https://space.bilibili.com/206214
const mod = 1_000_000_007
const mx = 500_000

var fac [mx]int  // fac[i] = i!
var invF [mx]int // invF[i] = i!^-1 = pow(i!, mod-2)

func init() {
	fac[0] = 1
	for i := 1; i < mx; i++ {
		fac[i] = fac[i-1] * i % mod
	}

	invF[mx-1] = pow(fac[mx-1], mod-2)
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

// 从 n 个数中选 m 个数的方案数
func comb(n, m int) int {
	return fac[n] * invF[m] % mod * invF[n-m] % mod
}

func countValidSequences(n, k int) int {
	ans := comb(n-1, k-1)
	if (n-k)%2 == 0 {
		ans = (ans - comb((n+k)/2-1, k-1) + mod) % mod // +mod 保证答案非负
	}
	return ans
}
