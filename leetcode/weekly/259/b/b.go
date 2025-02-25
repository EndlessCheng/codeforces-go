package main

// github.com/EndlessCheng/codeforces-go
func sumOfBeauties(nums []int) (ans int) {
	n := len(nums)
	sufMin := make([]int, n) // 后缀最小值
	sufMin[n-1] = nums[n-1]
	for i := n - 2; i > 1; i-- {
		sufMin[i] = min(sufMin[i+1], nums[i])
	}

	preMax := nums[0] // 前缀最大值
	for i := 1; i < n-1; i++ {
		x := nums[i]
		// 此时 preMax 表示 [0,i-1] 中的最大值
		if preMax < x && x < sufMin[i+1] {
			ans += 2
		} else if nums[i-1] < x && x < nums[i+1] {
			ans++
		}
		// 更新后 preMax 表示 [0,i] 中的最大值
		preMax = max(preMax, x)
	}
	return
}
