package main

import (
	"math"
	"slices"
)

// https://space.bilibili.com/206214
const mx = 1_000_001
var divisors [mx][]int32

func init() {
	// 本题 k > 1
	for i := int32(2); i < mx; i++ {
		for j := i; j < mx; j += i { // 枚举 i 的倍数 j
			divisors[j] = append(divisors[j], i) // i 是 j 的因子
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
	// 收集所有因子
	allDivisors := []int32{}
	for _, x := range nums {
		allDivisors = append(allDivisors, divisors[x]...)
	}

	if len(allDivisors) == 0 {
		// 每个数都是 1
		// 最优是只选一个 1（分数差为 -1），最小 k 为 2
		return mod - 2
	}

	// 排序去重
	slices.Sort(allDivisors)
	allDivisors = slices.Compact(allDivisors)

	maxDiff, bestK := math.MinInt, 0
	// 枚举因子作为 k，计算最大子数组和
	for _, d := range allDivisors {
		k := int(d)
		diff := maxSubArray(nums, k)
		if diff > maxDiff {
			maxDiff, bestK = diff, k
		}
	}

	// 保证结果非负
	return (maxDiff*bestK%mod + mod) % mod
}
