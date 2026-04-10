package main

import "math"

// https://space.bilibili.com/206214
func minimumDistance1(nums []int) int {
	pos := map[int][]int{}
	for i, x := range nums {
		if i >= 2 && x == nums[i-1] && x == nums[i-2] {
			return 4
		}
		pos[x] = append(pos[x], i)
	}

	ans := math.MaxInt
	for _, p := range pos {
		for i := 2; i < len(p); i++ {
			ans = min(ans, (p[i]-p[i-2])*2)
		}
	}

	if ans == math.MaxInt {
		return -1
	}
	return ans
}

func minimumDistance(nums []int) int {
	n := len(nums)
	last := make([]int, n+1)
	last2 := make([]int, n+1)
	for i := range last {
		last[i] = -n
		last2[i] = -n // i-last2[x] 不会把 ans 变小
	}

	ans := n
	for i, x := range nums {
		ans = min(ans, i-last2[x])
		last2[x] = last[x]
		last[x] = i
	}

	if ans == n {
		return -1
	}
	return ans * 2
}
