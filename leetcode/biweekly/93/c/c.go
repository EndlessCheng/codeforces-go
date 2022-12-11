package main

// https://space.bilibili.com/206214
func maxJump(stones []int) int {
	ans := stones[1] - stones[0]
	for i := 2; i < len(stones); i++ {
		ans = max(ans, stones[i]-stones[i-2])
	}
	return ans
}

func max(a, b int) int { if b > a { return b }; return a }
