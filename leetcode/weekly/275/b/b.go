package main

// github.com/EndlessCheng/codeforces-go
func minSwaps(nums []int) int {
	// 统计 1 的个数
	k := 0
	for _, x := range nums {
		k += x
	}
	if k == 0 { // 没有 1，无需交换
		return 0
	}

	n := len(nums)
	max1, cnt1 := 0, 0
	for i := range n + k - 1 {
		// 1. 进入窗口
		cnt1 += nums[i%n]
		if i < k-1 { // 窗口大小不足 k
			continue
		}
		// 2. 更新答案
		max1 = max(max1, cnt1)
		// 3. 离开窗口，为下一个循环做准备
		cnt1 -= nums[i-k+1] // 由于我们保证 i < n+k-1，所以 i-k+1 < n
	}
	return k - max1
}
