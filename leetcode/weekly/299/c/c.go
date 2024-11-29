package main

// https://space.bilibili.com/206214/dynamic
func solve(nums1, nums2 []int) int {
	var s1, maxSum, f int
	for i, x := range nums1 {
		s1 += x
		f = max(f, 0) + nums2[i] - x
		maxSum = max(maxSum, f)
	}
	return s1 + maxSum
}

func maximumsSplicedArray(nums1, nums2 []int) int {
	return max(solve(nums1, nums2), solve(nums2, nums1))
}
