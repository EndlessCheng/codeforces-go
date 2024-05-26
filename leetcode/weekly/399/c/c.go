package main

import "slices"

// https://space.bilibili.com/206214
func numberOfPairs(nums1, nums2 []int, k int) (ans int64) {
	cnt1 := map[int]int{}
	for _, x := range nums1 {
		if x%k == 0 {
			cnt1[x/k]++
		}
	}
	cnt2 := map[int]int{}
	for _, x := range nums2 {
		cnt2[x]++
	}

	m := slices.Max(nums1) / k
	for i, c := range cnt2 {
		s := 0
		for j := i; j <= m; j += i {
			s += cnt1[j]
		}
		ans += int64(s * c)
	}
	return
}
