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
	if len(cnt1) == 0 {
		return
	}

	cnt2 := map[int]int{}
	for _, x := range nums2 {
		cnt2[x]++
	}

	u := slices.Max(nums1) / k
	for x, cnt := range cnt2 {
		s := 0
		for y := x; y <= u; y += x { // 枚举倍数
			s += cnt1[y]
		}
		ans += int64(s * cnt)
	}
	return
}
