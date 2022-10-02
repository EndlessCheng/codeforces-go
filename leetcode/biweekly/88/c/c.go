package main

// https://space.bilibili.com/206214
func xorAllNums(nums1, nums2 []int) (ans int) {
	if len(nums2)%2 > 0 {
		for _, x := range nums1 {
			ans ^= x
		}
	}
	if len(nums1)%2 > 0 {
		for _, x := range nums2 {
			ans ^= x
		}
	}
	return
}
