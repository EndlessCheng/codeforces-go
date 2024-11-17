package main

// https://space.bilibili.com/206214
func isZeroArray(nums []int, queries [][]int) bool {
	diff := make([]int, len(nums)+1)
	for _, q := range queries {
		// 区间 [l,r] 中的数都加一
		diff[q[0]]++
		diff[q[1]+1]--
	}

	sumD := 0
	for i, x := range nums {
		sumD += diff[i]
		// 此时 sumD 表示 x=nums[i] 要减掉多少
		if x > sumD { // x 无法变成 0
			return false
		}
	}
	return true
}
