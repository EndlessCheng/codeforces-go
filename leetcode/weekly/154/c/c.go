package main

// https://space.bilibili.com/206214
// 53. 最大子数组和（遍历 nums repeat 次）
// 本题允许子数组为空，ans 可以初始化成 0
func maxSubArray(nums []int, repeat int) (ans int) {
	f := 0
	for range repeat {
		for _, x := range nums {
			f = max(f, 0) + x
			ans = max(ans, f)
		}
	}
	return
}

func kConcatenationMaxSum(arr []int, k int) int {
	if k == 1 {
		return maxSubArray(arr, 1) // arr 的最大子数组和
	}
	ans := maxSubArray(arr, 2) // arr+arr 的最大子数组和
	s := 0
	for _, x := range arr {
		s += x
	}
	ans += max(s, 0) * (k - 2)
	return ans % 1_000_000_007
}
