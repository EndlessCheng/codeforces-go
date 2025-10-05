package main

// https://space.bilibili.com/206214
func longestSubsequence(nums []int) int {
	sum, xor := 0, 0
	for _, x := range nums {
		sum += x
		xor ^= x
	}
	if sum == 0 {
		return 0 // nums 全为 0，无解
	}

	ans := len(nums)
	if xor == 0 {
		ans-- // 去掉 nums 的一个非零元素，就可以使 xor 不为零
	}
	return ans
}
