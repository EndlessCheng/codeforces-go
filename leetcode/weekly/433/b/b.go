package main

import "slices"

// https://space.bilibili.com/206214
const mod = 1_000_000_007
const mx = 100_000

var f [mx]int    // f[i] = i!
var invF [mx]int // invF[i] = i!^-1

func init() {
	f[0] = 1
	for i := 1; i < mx; i++ {
		f[i] = f[i-1] * i % mod
	}

	invF[mx-1] = pow(f[mx-1], mod-2)
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
	if m > n {
		return 0
	}
	return f[n] * invF[m] % mod * invF[n-m] % mod
}

func minMaxSums(nums []int, k int) (ans int) {
	slices.Sort(nums)
	s := 1
	for i, x := range nums {
		ans = (ans + s*(x+nums[len(nums)-1-i])) % mod
		s = (s*2 - comb(i, k-1) + mod) % mod
	}
	return
}
