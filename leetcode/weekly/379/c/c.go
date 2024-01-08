package main

// https://space.bilibili.com/206214
func maximumSetSize(nums1, nums2 []int) int {
	set1 := map[int]bool{}
	for _, x := range nums1 {
		set1[x] = true
	}
	set2 := map[int]bool{}
	common := 0
	for _, x := range nums2 {
		if set2[x] {
			continue
		}
		set2[x] = true
		if set1[x] {
			common++
		}
	}

	n := len(nums1)
	c1 := min(len(set1)-common, n/2)
	c2 := min(len(set2)-common, n/2)
	return min(n, c1+c2+common)
}
