package main

// https://space.bilibili.com/206214
func maximumOr(nums []int, k int) int64 {
	n := len(nums)
	suf := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		suf[i] = suf[i+1] | nums[i]
	}
	ans, pre := 0, 0
	for i, x := range nums {
		ans = max(ans, pre|x<<k|suf[i+1])
		pre |= x
	}
	return int64(ans)
}

func max(a, b int) int { if a < b { return b }; return a }
