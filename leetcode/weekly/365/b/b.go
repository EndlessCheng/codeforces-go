package main

// https://space.bilibili.com/206214
func maximumTripletValue(nums []int) int64 {
	ans, maxDiff, preMax := 0, 0, 0
	for _, x := range nums {
		ans = max(ans, maxDiff*x)
		maxDiff = max(maxDiff, preMax-x)
		preMax = max(preMax, x)
	}
	return int64(ans)
}

func max(a, b int) int { if b > a { return b }; return a }
