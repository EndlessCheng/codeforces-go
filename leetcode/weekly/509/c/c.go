package main

import (
	"math"
)

// https://space.bilibili.com/206214
const mx = 1_000_001

var primeDivisors [mx][]int32

// 预处理每个数的质因子
func init() {
	for i := int32(2); i < mx; i++ {
		if primeDivisors[i] == nil { // i 是质数
			for j := i; j < mx; j += i { // 枚举 i 的倍数 j
				primeDivisors[j] = append(primeDivisors[j], i) // i 是 j 的质因子
			}
		}
	}
}

func divisibleGame(nums []int) (ans int) {
	const mod = 1_000_000_007

	n := len(nums)
	sum := make([]int, n+1)
	for i, x := range nums {
		sum[i+1] = sum[i] + x
	}
	if sum[n] == n {
		// 每个数都是 1
		// 最优是只选一个 1（分数差为 -1），最小 k 为 2
		return mod - 2
	}

	f := map[int32]int{}
	last := map[int32]int{}
	maxDiff, bestK := math.MinInt, int32(0)

	for i, x := range nums {
		for _, p := range primeDivisors[x] {
			diff := max(f[p]-sum[i]+sum[last[p]], 0) + x
			f[p] = diff
			if diff > maxDiff || diff == maxDiff && p < bestK {
				maxDiff, bestK = diff, p
			}
			last[p] = i + 1
		}
	}

	return maxDiff * int(bestK) % mod
}
