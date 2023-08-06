package main

import "sort"

// https://space.bilibili.com/206214
func minimumTime(nums1, nums2 []int, x int) int {
	s1, s2, n := 0, 0, len(nums1)
	id := make([]int, n)
	for i := range id {
		id[i] = i
		s1 += nums1[i]
		s2 += nums2[i]
	}
	sort.Slice(id, func(i, j int) bool { return nums2[id[i]] < nums2[id[j]] })

	f := make([]int, n+1)
	for _, i := range id {
		for j := n; j > 0; j-- {
			f[j] = max(f[j], f[j-1]+nums1[i]+nums2[i]*j)
		}
	}

	for t, v := range f {
		if s1+s2*t-v <= x {
			return t
		}
	}
	return -1
}

func max(a, b int) int { if b > a { return b }; return a }