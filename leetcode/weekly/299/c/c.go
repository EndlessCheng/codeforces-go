package main

// https://space.bilibili.com/206214/dynamic
func solve(nums1, nums2 []int) int {
	s1, maxSum, s := 0, 0, 0
	for i, x := range nums1 {
		s1 += x
		s = max(s+nums2[i]-x, 0)
		maxSum = max(maxSum, s)
	}
	return s1 + maxSum
}

func maximumsSplicedArray(nums1, nums2 []int) int {
	return max(solve(nums1, nums2), solve(nums2, nums1))
}

func max(a, b int) int { if b > a { return b }; return a }
