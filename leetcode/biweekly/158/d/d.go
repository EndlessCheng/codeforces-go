package main

import "slices"

// https://space.bilibili.com/206214
func goodSubtreeSum(vals, par []int) (ans int) {
	const mod = 1_000_000_007
	const D = 10
	n := len(par)
	g := make([][]int, n)
	for i := 1; i < n; i++ {
		p := par[i]
		g[p] = append(g[p], i)
	}

	var dfs func(int) [1 << D]int
	dfs = func(x int) (f [1 << D]int) {
		// 计算 vals[x] 的数位集合 mask
		mask := 0
		for v := vals[x]; v > 0; v /= D {
			d := v % D
			if mask>>d&1 > 0 { // d 在集合 mask 中
				mask = 0 // 不符合要求
				break
			}
			mask |= 1 << d // 把 d 加到集合 mask 中
		}

		if mask > 0 {
			f[mask] = vals[x]
		}

		// 同一个集合 i 至多选一个，直接取 max
		for _, y := range g[x] {
			fy := dfs(y)
			for i, sum := range fy {
				f[i] = max(f[i], sum)
			}
		}

		for i := range f {
			// 枚举集合 i 的非空真子集
			for sub := i & (i - 1); sub > 0; sub = (sub - 1) & i {
				f[i] = max(f[i], f[sub]+f[i^sub])
			}
		}

		ans += slices.Max(f[:])
		return
	}
	dfs(0)
	return ans % mod
}
