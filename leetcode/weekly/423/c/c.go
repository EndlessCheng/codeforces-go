package main

import "slices"

// https://space.bilibili.com/206214
func sumOfGoodSubsequences(nums []int) (ans int) {
	const mod = 1_000_000_007
	mx := slices.Max(nums)
	f := make([]int, mx+3)
	cnt := make([]int, mx+3)
	for _, x := range nums {
		// 为避免出现 -1，所有下标加一
		c := cnt[x] + cnt[x+2] + 1
		f[x+1] = (f[x] + f[x+1] + f[x+2] + x*c) % mod
		cnt[x+1] = (cnt[x+1] + c) % mod
	}

	for _, s := range f {
		ans += s
	}
	return ans % mod
}

func sumOfGoodSubsequences2(nums []int) (ans int) {
	const mod = 1_000_000_007
	f := map[int]int{}
	cnt := map[int]int{}
	for _, x := range nums {
		c := cnt[x-1] + cnt[x+1] + 1
		f[x] = (f[x] + f[x-1] + f[x+1] + x*c) % mod
		cnt[x] = (cnt[x] + c) % mod
	}

	for _, s := range f {
		ans += s
	}
	return ans % mod
}
