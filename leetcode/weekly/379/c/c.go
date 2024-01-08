package main

// https://space.bilibili.com/206214
func maximumSetSize(nums1, nums2 []int) int {
	set1 := map[int]bool{}
	for _, x := range nums1 {
		set1[x] = true
	}
	all := len(set1)
	set2 := map[int]bool{}
	for _, x := range nums2 {
		if set2[x] {
			continue
		}
		set2[x] = true
		if !set1[x] {
			all++
		}
	}

	n := len(nums1)
	c1 := min(len(set1), n/2)
	c2 := min(len(set2), n/2)
	return min(all, c1+c2)
}
