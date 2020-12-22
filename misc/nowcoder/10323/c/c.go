package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func city(n, k int, a []int) []int {
	type viPair struct{ v, i int }
	vi := make([]viPair, n)
	for i, v := range a {
		vi[i] = viPair{v, i}
	}
	sort.Slice(vi, func(i, j int) bool { return vi[i].v < vi[j].v })

	const mod int = 1e9 + 7
	const mx int = 2e5
	F := [mx + 1]int{1}
	for i := 1; i <= mx; i++ {
		F[i] = F[i-1] * i % mod
	}
	pow := func(x, n int) (res int) {
		res = 1
		for ; n > 0; n >>= 1 {
			if n&1 == 1 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return res
	}
	invF := [...]int{mx: pow(F[mx], mod-2)}
	for i := mx; i > 0; i-- {
		invF[i-1] = invF[i] * i % mod
	}
	C := func(n, k int) int { return F[n] * invF[k] % mod * invF[n-k] % mod }
	inv := func(a int) int { return pow(a, mod-2) }
	div := func(a, b int) int { return a * inv(b) % mod }

	all := C(n, k)
	ans := make([]int, n)
	for i := k; i <= n; i++ {
		ans[vi[i-1].i] = div(C(i-1, k-1), all)
	}
	return ans
}
