package main

// https://space.bilibili.com/206214
func countSubarrays(nums []int, minK, maxK int) (ans int64) {
	minI, maxI, i0 := -1, -1, -1
	for i, x := range nums {
		if x == minK {
			minI = i
		}
		if x == maxK {
			maxI = i
		}
		if x < minK || x > maxK {
			i0 = i
		}
		ans += int64(max(min(minI, maxI)-i0, 0))
	}
	return
}

func min(a, b int) int { if b < a { return b }; return a }
func max(a, b int) int { if b > a { return b }; return a }
