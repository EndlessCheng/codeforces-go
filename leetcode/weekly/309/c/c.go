package main

// https://space.bilibili.com/206214
func longestNiceSubarray(nums []int) (ans int) {
	left, or := 0, 0
	for right, x := range nums {
		for or&x > 0 {
			or ^= nums[left]
			left += 1
		}
		or |= x
		ans = max(ans, right-left+1)
	}
	return
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
