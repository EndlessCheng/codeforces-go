package main

// https://space.bilibili.com/206214
func minOperations(nums1, nums2 []int) int {
	n := len(nums1)
	f := func(last1, last2 int) (res int) {
		for i, x := range nums1[:n-1] {
			y := nums2[i]
			if x > last1 || y > last2 {
				if y > last1 || x > last2 {
					return n + 1
				}
				res++
			}
		}
		return
	}
	ans := min(f(nums1[n-1], nums2[n-1]), 1+f(nums2[n-1], nums1[n-1]))
	if ans > n {
		return -1
	}
	return ans
}
