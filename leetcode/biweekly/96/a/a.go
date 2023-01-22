package main

// https://space.bilibili.com/206214
func getCommon(nums1, nums2 []int) int {
	j, m := 0, len(nums2)
	for _, x := range nums1 {
		for j < m && nums2[j] < x {
			j++
		}
		if j < m && nums2[j] == x {
			return x
		}
	}
	return -1
}
