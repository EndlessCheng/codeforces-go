package main

import (
	"slices"
	"sort"
)

// https://space.bilibili.com/206214
var pal = make([]int, 0, 109999)

func init() {
	// 按顺序从小到大生成所有回文数
	for base := 1; base <= 10000; base *= 10 {
		for i := base; i < base*10; i++ {
			x := i
			for t := i / 10; t > 0; t /= 10 {
				x = x*10 + t%10
			}
			pal = append(pal, x)
		}
		if base <= 1000 {
			for i := base; i < base*10; i++ {
				x := i
				for t := i; t > 0; t /= 10 {
					x = x*10 + t%10
				}
				pal = append(pal, x)
			}
		}
	}
	pal = append(pal, 1e9+1) // 哨兵，防止下标越界
}

func minimumCost(nums []int) int64 {
	// 注：排序只是为了找中位数，如果用快速选择算法，可以做到 O(n)
	slices.Sort(nums)

	// 返回所有 nums[i] 变成 pal[i] 的总代价
	cost := func(i int) (res int64) {
		target := pal[i]
		for _, x := range nums {
			res += int64(abs(x - target))
		}
		return
	}

	n := len(nums)
	i := sort.SearchInts(pal, nums[(n-1)/2])
	if pal[i] <= nums[n/2] { // 回文数在中位数范围内
		return cost(i) // 直接变成 pal[i]
	}
	return min(cost(i-1), cost(i)) // 枚举离中位数最近的两个回文数 pal[i-1] 和 pal[i]
}

func abs(x int) int { if x < 0 { return -x }; return x }
