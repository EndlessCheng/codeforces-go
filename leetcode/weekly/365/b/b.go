package main

// https://space.bilibili.com/206214
func maximumTripletValue(nums []int) int64 {
	var ans, maxDiff, preMax int
	for _, x := range nums {
		ans = max(ans, maxDiff*x)
		maxDiff = max(maxDiff, preMax-x)
		preMax = max(preMax, x)
	}
	return int64(ans)
}
