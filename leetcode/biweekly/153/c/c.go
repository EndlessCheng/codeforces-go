package main

import "math"

// https://space.bilibili.com/206214
func minimumCost(nums, cost []int, k int) int64 {
	n := len(nums)
	s := make([]int, n+1)
	for i, c := range cost {
		s[i+1] = s[i] + c
	}

	f := make([]int, n+1)
	sumNum := 0
	for i, x := range nums {
		i++
		sumNum += x
		res := math.MaxInt
		for j := range i {
			res = min(res, f[j]+sumNum*(s[i]-s[j])+k*(s[n]-s[j]))
		}
		f[i] = res
	}
	return int64(f[n])
}
