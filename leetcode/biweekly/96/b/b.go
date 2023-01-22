package main

// https://space.bilibili.com/206214
func minOperations(nums1, nums2 []int, k int) (ans int64) {
	sum := 0
	for i, x := range nums1 {
		x -= nums2[i]
		if k > 0 {
			if x%k != 0 {
				return -1
			}
			sum += x / k
			if x > 0 {
				ans += int64(x / k)
			}
		} else if x != 0 {
			return -1
		}
	}
	if k > 0 && sum != 0 {
		return -1
	}
	return
}
