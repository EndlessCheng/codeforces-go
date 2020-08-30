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
	g := make([][2]int, len(a)+1)
	root := a[0]
	var v int
	var put func(int)
	put = func(o int) {
		d := 0
		if v > o {
			d = 1
		}
		if g[o][d] == 0 {
			g[o][d] = v
		} else {
			put(g[o][d])
		}
	}
	for _, v = range a[1:] {
		put(root)
	}

	ans := 1
	var f func(int) int
	f = func(v int) int {
		if v == 0 {
			return 0
		}
		l, r := f(g[v][0]), f(g[v][1])
		if l > 0 && r > 0 {
			ans = ans * F[l+r] % mod * invF[l] % mod * invF[r] % mod
		}
		return 1 + l + r
	}
	f(root)
	return (ans + mod - 1) % mod
}
