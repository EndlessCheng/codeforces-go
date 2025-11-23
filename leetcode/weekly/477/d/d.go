package main

import "math/bits"

// https://space.bilibili.com/206214
const mod = 1_000_000_007
const maxN = 100_001

var pow2 = [maxN]int{1}

func init() {
	// 预处理 2 的幂
	for i := 1; i < maxN; i++ {
		pow2[i] = pow2[i-1] * 2 % mod
	}
}

func countEffective(nums []int) int {
	or := 0
	for _, x := range nums {
		or |= x
	}

	mx := bits.Len(uint(or))
	f := make([]int, 1<<mx)
	for _, x := range nums {
		f[x]++
	}
	for i := range mx {
		for s := 0; s < 1<<mx; s++ {
			s |= 1 << i
			f[s] += f[s^1<<i]
		}
	}
	// 计算完毕后，f[s] 表示 nums 中的是 s 的子集的元素个数

	ans := pow2[len(nums)] // 所有子序列的个数
	// 枚举 or 的所有子集（包括空集）
	for sub, ok := or, true; ok; ok = sub != or {
		sign := 1 - bits.OnesCount(uint(or^sub))%2*2
		ans -= sign * pow2[f[sub]]
		sub = (sub - 1) & or
	}
	return (ans%mod + mod) % mod // 保证结果非负
}
