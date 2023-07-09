package main

// https://space.bilibili.com/206214
func maxNonDecreasingLength(nums1, nums2 []int) int {
	ans, n := 1, len(nums1)
	f0, f1 := 1, 1
	for i := 1; i < n; i++ {
		f, g := 1, 1
		if nums1[i-1] <= nums1[i] {
			f = f0 + 1
		}
		if nums2[i-1] <= nums1[i] {
			f = max(f, f1+1)
		}
		if nums1[i-1] <= nums2[i] {
			g = f0 + 1
		}
		if nums2[i-1] <= nums2[i] {
			g = max(g, f1+1)
		}
		f0, f1 = f, g
		ans = max(ans, max(f0, f1))
	}
	return ans
}

func max(a, b int) int { if b > a { return b }; return a }
