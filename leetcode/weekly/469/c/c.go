package main

import "slices"

// https://space.bilibili.com/206214
func zigZagArrays1(n, l, r int) (ans int) {
	const mod = 1_000_000_007
	k := r - l + 1
	f0 := make([]int, k) // 后两个数递增
	f1 := make([]int, k) // 后两个数递减
	for i := range f0 {
		f0[i] = 1
		f1[i] = 1
	}

	s0 := make([]int, k+1)
	s1 := make([]int, k+1)
	for range n - 1 {
		for j, v := range f0 {
			s0[j+1] = s0[j] + v
			s1[j+1] = s1[j] + f1[j]
		}
		for j := range f0 {
			f0[j] = s1[j] % mod
			f1[j] = (s0[k] - s0[j+1]) % mod
		}
	}

	for j, v := range f0 {
		ans += v + f1[j]
	}
	return ans % mod
}

func zigZagArrays(n, l, r int) (ans int) {
	const mod = 1_000_000_007
	k := r - l + 1
	f := make([]int, k)
	for i := range f {
		f[i] = 1
	}

	for i := 1; i < n; i++ {
		pre := 0
		for j, v := range f {
			f[j] = pre % mod
			pre += v
		}
		slices.Reverse(f)
	}

	for _, v := range f {
		ans += v
	}
	return ans * 2 % mod
}
