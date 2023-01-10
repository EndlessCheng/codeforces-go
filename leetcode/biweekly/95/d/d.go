package main

import (
	"math"
	"sort"
)

// https://space.bilibili.com/206214
func maxPower(stations []int, r int, k int) int64 {
	n := len(stations)
	sum := make([]int, n+1) // 前缀和
	for i, x := range stations {
		sum[i+1] = sum[i] + x
	}
	mn := math.MaxInt
	for i := range stations {
		stations[i] = sum[min(i+r+1, n)] - sum[max(i-r, 0)] // 电量
		mn = min(mn, stations[i])
	}
	return int64(mn + sort.Search(k, func(minPower int) bool {
		minPower += mn + 1 // 改为二分最小的不满足要求的值，这样 sort.Search 返回的就是最大的满足要求的值
		diff := make([]int, n) // 差分数组
		sumD, need := 0, 0
		for i, power := range stations {
			sumD += diff[i] // 累加差分值
			m := minPower - power - sumD
			if m > 0 { // 需要 m 个供电站
				need += m
				if need > k { // 提前退出这样快一些
					return true // 不满足要求
				}
				sumD += m // 差分更新
				if i+r*2+1 < n {
					diff[i+r*2+1] -= m // 差分更新
				}
			}
		}
		return false
	}))
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}
func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
