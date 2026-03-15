package main

// https://space.bilibili.com/206214
func minCost(nums1, nums2 []int) (ans int) {
	diff := map[int]int{}
	for _, x := range nums1 {
		diff[x]++
	}
	for _, x := range nums2 {
		diff[x]--
	}

	for _, d := range diff {
		if d%2 != 0 {
			return -1
		}
		if d > 0 {
			ans += d
		}
	}
	return ans / 2
}
