package main

// https://space.bilibili.com/206214
func minIncrementOperations(nums []int, k int) int64 {
	var f0, f1, f2 int
	for _, x := range nums {
		inc := f0 + max(k-x, 0)
		f0 = min(inc, f1)
		f1 = min(inc, f2)
		f2 = inc
	}
	return int64(f0)
}

func min(a, b int) int { if b < a { return b }; return a }
func max(a, b int) int { if b > a { return b }; return a }
