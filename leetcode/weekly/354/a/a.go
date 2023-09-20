package main

// https://space.bilibili.com/206214
func sumOfSquares(nums []int) (ans int) {
	n := len(nums)
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			ans += nums[i-1] * nums[i-1] // 注意数组的下标还是从 0 开始的
			if i*i < n {
				ans += nums[n/i-1] * nums[n/i-1]
			}
		}
	}
	return
}
