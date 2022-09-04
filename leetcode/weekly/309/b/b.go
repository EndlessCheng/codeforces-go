package main

// https://space.bilibili.com/206214
func numberOfWays(startPos, endPos, k int) int {
	type pair struct{ x, y int }
	dp := map[pair]int{}
	var f func(int, int) int
	f = func(x, left int) int {
		if abs(x-endPos) > left {
			return 0
		}
		if left == 0 {
			return 1
		}
		p := pair{x, left}
		if v, ok := dp[p]; ok {
			return v
		}
		res := (f(x-1, left-1) + f(x+1, left-1)) % (1e9 + 7)
		dp[p] = res
		return res
	}
	return f(startPos, k)
}
func abs(x int) int { if x < 0 { return -x }; return x }

