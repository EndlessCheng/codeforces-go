package main

// https://space.bilibili.com/206214
func isArraySpecial(nums []int, queries [][]int) []bool {
	n := len(nums)
	lastSame := make([]int, n)
	for i := 1; i < n; i++ {
		if nums[i-1]%2 == nums[i]%2 {
			lastSame[i] = i
		} else {
			lastSame[i] = lastSame[i-1]
		}
	}
	ans := make([]bool, len(queries))
	for i, q := range queries {
		ans[i] = lastSame[q[1]] <= q[0]
	}
	return ans
}

func isArraySpecial2(nums []int, queries [][]int) []bool {
	s := make([]int, len(nums))
	for i := 1; i < len(nums); i++ {
		s[i] = s[i-1] + (nums[i]^nums[i-1]^1)&1
	}
	ans := make([]bool, len(queries))
	for i, q := range queries {
		ans[i] = s[q[0]] == s[q[1]]
	}
	return ans
}
