package main

// https://space.bilibili.com/206214
func findIntersectionValues(nums1, nums2 []int) []int {
	set1 := map[int]int{}
	for _, x := range nums1 {
		set1[x] = 1
	}
	set2 := map[int]int{}
	for _, x := range nums2 {
		set2[x] = 1
	}
	
	ans := [2]int{}
	for _, x := range nums1 {
		ans[0] += set2[x]
	}
	for _, x := range nums2 {
		ans[1] += set1[x]
	}
	return ans[:]
}
