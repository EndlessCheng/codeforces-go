package main

// https://space.bilibili.com/206214
func maximumSetSize(nums1, nums2 []int) int {
	set1 := map[int]bool{}
	for _, x := range nums1 {
		set1[x] = true
	}
	set2 := map[int]bool{}
	for _, x := range nums2 {
		set2[x] = true
	}

	common := 0
	for x := range set1 {
		if set2[x] {
			common++
		}
	}

	n := len(nums1)
	n1 := min(len(set1)-common, n/2) // 不考虑交集，nums1 至多保留 n1 个数
	n2 := min(len(set2)-common, n/2) // 不考虑交集，nums2 至多保留 n2 个数
	return min(n1+n2+common, n)
}
