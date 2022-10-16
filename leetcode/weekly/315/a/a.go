package main

// https://space.bilibili.com/206214
func findMaxK(nums []int) int {
	ans := -1
	has := map[int]bool{}
	for _, x := range nums {
		if abs(x) > ans && has[-x] {
			ans = abs(x)
		}
		has[x] = true
	}
	return ans
}

func abs(x int) int { if x < 0 { return -x }; return x }
