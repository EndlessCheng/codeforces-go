package main

import (
	"math"
	"slices"
)

// https://space.bilibili.com/206214
const mx = 1_000_001
var primeDivisors [mx][]int32

func init() {
	for i := int32(2); i < mx; i++ {
		if primeDivisors[i] == nil { // i 是质数
			for j := i; j < mx; j += i { // 枚举 i 的倍数 j
				primeDivisors[j] = append(primeDivisors[j], i) // i 是 j 的质因子
			}
		}
	}
}

// 53. 最大子数组和（如果 nums[i] 不是 k 的倍数，则视作 -nums[i]）
func maxSubArray(nums []int, k int) int {
	ans := math.MinInt
	f := 0
	for _, x := range nums {
		if x%k != 0 {
			x = -x
		}
		f = max(f, 0) + x
		ans = max(ans, f)
	}
	return ans
}

func divisibleGame(nums []int) (ans int) {
	const mod = 1_000_000_007
	// 收集所有质因子
	allPrimeDivisors := []int32{}
	for _, x := range nums {
		allPrimeDivisors = append(allPrimeDivisors, primeDivisors[x]...)
	}

	if len(allPrimeDivisors) == 0 {
		// 每个数都是 1
		// 最优是只选一个 1（分数差为 -1），最小 k 为 2
		return mod - 2
	}

	// 排序去重
	slices.Sort(allPrimeDivisors)
	allPrimeDivisors = slices.Compact(allPrimeDivisors)

	maxDiff, bestK := math.MinInt, 0
	// 枚举质因子作为 k，计算最大子数组和
	for _, d := range allPrimeDivisors {
		k := int(d)
		diff := maxSubArray(nums, k)
		if diff > maxDiff {
			maxDiff, bestK = diff, k
		}
	}

	return maxDiff * bestK % mod
}
