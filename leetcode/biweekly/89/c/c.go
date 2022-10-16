package main

// https://space.bilibili.com/206214
func minimizeArrayValue(nums []int) (ans int) {
	s := 0
	for i, x := range nums {
		s += x
		ans = max(ans, (s+i)/(i+1))
	}
	return
}

func max(a, b int) int { if b > a { return b }; return a }
