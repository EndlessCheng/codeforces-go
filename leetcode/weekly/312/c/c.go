package main

// https://space.bilibili.com/206214
func goodIndices(nums []int, k int) (ans []int) {
	n := len(nums)
	dec := make([]int, n)
	dec[n-1] = 1
	for i := n - 2; i > k; i-- {
		if nums[i] <= nums[i+1] {
			dec[i] = dec[i+1] + 1 // 递推
		} else {
			dec[i] = 1
		}
	}
	for i, inc := 1, 1; i < n-k; i++ {
		if inc >= k && dec[i+1] >= k {
			ans = append(ans, i)
		}
		if nums[i-1] >= nums[i] {
			inc++ // 递推
		} else {
			inc = 1
		}
	}
	return
}
