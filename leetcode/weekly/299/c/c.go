package main

// https://space.bilibili.com/206214/dynamic
func maxSubarraySum(a []int) int {
	maxS, s := 0, 0
	for _, v := range a {
		s = max(s+v, 0)
		maxS = max(maxS, s)
	}
	return maxS
}

func maximumsSplicedArray(nums1, nums2 []int) (ans int) {
	n := len(nums1)
	f := func(a, b []int) {
		diff := make([]int, n)
		for i, v := range b {
			diff[i] = v - a[i]
		}
		sum := 0
		for _, v := range a {
			sum += v
		}
		sum += maxSubarraySum(diff)
		ans = max(ans, sum)
	}
	f(nums1, nums2)
	f(nums2, nums1)
	return
}

func max(a, b int) int { if b > a { return b }; return a }
