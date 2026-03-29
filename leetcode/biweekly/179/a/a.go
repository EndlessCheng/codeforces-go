package main

// https://space.bilibili.com/206214
func minAbsoluteDifference(nums []int) int {
	n := len(nums)
	ans := n
	// last[x] 表示 x+1 上一次出现的位置
	last := [2]int{-n, -n} // i - (-n) >= n，不会让 ans 变小

	for i, x := range nums {
		if x > 0 {
			// 如果 x 是 1，那么找上一个 2 的位置
			// 如果 x 是 2，那么找上一个 1 的位置
			x--
			ans = min(ans, i-last[x^1])
			last[x] = i
		}
	}

	if ans == n {
		return -1
	}
	return ans
}
