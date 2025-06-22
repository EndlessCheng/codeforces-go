package main

import "math"

// https://space.bilibili.com/206214
func minSwaps(nums []int) int {
	pos1 := []int{}
	for i, x := range nums {
		if x%2 != 0 {
			pos1 = append(pos1, i)
		}
	}

	n := len(nums)
	m := len(pos1)
	if abs(n-m*2) > 1 {
		return -1
	}

	// start=0 表示车去偶数下标，start=1 表示车去奇数下标
	calc := func(start int) (res int) {
		if (n-start+1)/2 != m {
			return math.MaxInt
		}
		for i, j := range pos1 {
			res += abs(i*2 + start - j)
		}
		return
	}
	return min(calc(0), calc(1))
}

func abs(x int) int { if x < 0 { return -x }; return x }
