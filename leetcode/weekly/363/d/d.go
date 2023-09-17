package main

// https://space.bilibili.com/206214
func maximumSum(nums []int) (ans int64) {
	n := len(nums)
	for i := 1; i <= n; i++ {
		sum := int64(0)
		for j := 1; i*j*j <= n; j++ {
			sum += int64(nums[i*j*j-1])
		}
		ans = max(ans, sum)
	}
	return
}

func max(a, b int64) int64 {
	if b > a {
		return b
	}
	return a
}
