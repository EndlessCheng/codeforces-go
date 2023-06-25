package main

// https://space.bilibili.com/206214
func numberOfGoodSubarraySplits(nums []int) int {
	const mod int = 1e9 + 7
	ans, pre := 1, -1
	for i, x := range nums {
		if x > 0 {
			if pre >= 0 {
				ans = ans * (i - pre) % mod
			}
			pre = i
		}
	}
	if pre < 0 { // 整个数组都是 0，没有好子数组
		return 0
	}
	return ans
}
