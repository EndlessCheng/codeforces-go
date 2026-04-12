package main

import (
	"cmp"
	"slices"
)

// https://space.bilibili.com/206214
const mod = 1_000_000_007
const mx = 10001

var pow2 = [mx]int{1}

func init() {
	// 预处理 2 的幂
	for i := 1; i < mx; i++ {
		pow2[i] = pow2[i-1] * 2 % mod
	}
}

func maxValue(nums1, nums0 []int) (ans int) {
	idx := make([]int, len(nums1))
	for i := range idx {
		idx[i] = i
	}
	slices.SortFunc(idx, func(i, j int) int {
		if nums0[i] == 0 {
			return -1
		}
		if nums0[j] == 0 {
			return 1
		}
		return cmp.Or(nums1[j]-nums1[i], nums0[i]-nums0[j])
	})

	for _, i := range idx {
		ans = ((ans+1)*pow2[nums1[i]] - 1) % mod * pow2[nums0[i]] % mod
	}
	return
}
