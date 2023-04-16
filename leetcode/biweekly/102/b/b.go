package main

// https://space.bilibili.com/206214
func findPrefixScore(nums []int) []int64 {
	ans := make([]int64, len(nums))
	mx, sum := 0, 0
	for i, x := range nums {
		mx = max(mx, x) // 前缀最大值
		sum += x + mx   // 累加前缀的得分
		ans[i] = int64(sum)
	}
	return ans
}

func max(a, b int) int { if a < b { return b }; return a }
