package main

// https://space.bilibili.com/206214
func longestSquareStreak(nums []int) (ans int) {
	set := map[int]bool{}
	for _, x := range nums {
		set[x] = true
	}
	dp := map[int]int{}
	var f func(int) int
	f = func(x int) int {
		if !set[x] {
			return 0
		}
		if v, ok := dp[x]; ok {
			return v
		}
		dp[x] = 1 + f(x*x)
		return dp[x]
	}
	for x := range set {
		ans = max(ans, f(x))
	}
	if ans == 1 {
		return -1
	}
	return
}

func max(a, b int) int { if b > a { return b }; return a }
