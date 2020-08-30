package main

// github.com/EndlessCheng/codeforces-go
const mx, mod int = 1e3, 1e9 + 7

var F, invF [mx + 1]int

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

func numOfWays(a []int) int {
	ans := 1
	var f func([]int) int
	f = func(a []int) int {
		if len(a) == 0 {
			return 0
		}
		var b, c []int
		for _, v := range a[1:] {
			if v < a[0] {
				b = append(b, v)
			} else {
				c = append(c, v)
			}
		}
		l, r := f(b), f(c)
		if l > 0 && r > 0 {
			ans = ans * F[l+r] % mod * invF[l] % mod * invF[r] % mod
		}
		return 1 + l + r
	}
	f(a)
	return (ans + mod - 1) % mod
}
