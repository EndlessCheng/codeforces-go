package main

// github.com/EndlessCheng/codeforces-go
const mx, mod int = 3e3, 1e9 + 7

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

func makeStringSorted(s string) (ans int) {
	cnt := make([]int, 26)
	for _, b := range s {
		cnt[b-'a']++
	}
	for i, b := range s {
		b -= 'a'
		rk := 0
		for _, v := range cnt[:b] {
			rk += v
		}
		m := rk * F[len(s)-1-i] % mod
		for _, c := range cnt {
			m = m * invF[c] % mod
		}
		ans += m
		cnt[b]--
	}
	return ans % mod
}

func pow(x, n int) int {
	res := 1
	for ; n > 0; n >>= 1 {
		if n&1 == 1 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return res
}
