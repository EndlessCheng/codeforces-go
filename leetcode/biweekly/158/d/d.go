package main

import (
	"maps"
	"slices"
)

// https://space.bilibili.com/206214
func goodSubtreeSum1(vals, par []int) (ans int) {
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

func goodSubtreeSumHSon(vals, par []int) (ans int) {
	const mod = 1_000_000_007
	const D = 10
	n := len(par)
	g := make([][]int, n)
	for i := 1; i < n; i++ {
		p := par[i]
		g[p] = append(g[p], i)
	}

	var init func(int) int
	init = func(x int) int {
		if g[x] == nil {
			return 1
		}
		size, hsz, hIdx := 1, 0, 0
		for i, y := range g[x] {
			sz := init(y)
			size += sz
			if sz > hsz {
				hsz, hIdx = sz, i
			}
		}
		// 把重儿子换到最前面
		g[x][0], g[x][hIdx] = g[x][hIdx], g[x][0]
		return size
	}
	init(0)

	type pair struct{ mask, val int }
	var dfs func(int) (map[int]int, []pair)
	dfs = func(x int) (f map[int]int, single []pair) {
		val := vals[x]

		// 计算 val 的数位集合 mask
		mask := 0
		for v := val; v > 0; v /= D {
			d := v % D
			if mask>>d&1 > 0 {
				mask = 0
				break
			}
			mask |= 1 << d
		}

		if g[x] == nil { // x 是叶子
			f = map[int]int{}
			if mask > 0 {
				ans += val
				f[mask] = val
				single = append(single, pair{mask, val})
			}
			return
		}

		f, single = dfs(g[x][0]) // 优先遍历重儿子
		update := func(msk, v int) {
			if v <= f[msk] {
				return
			}
			nf := maps.Clone(f)
			nf[msk] = v
			for msk2, s2 := range f {
				if msk&msk2 == 0 {
					nf[msk|msk2] = max(nf[msk|msk2], v+s2)
				}
			}
			f = nf
		}

		for _, y := range g[x][1:] {
			_, singleY := dfs(y)
			single = append(single, singleY...)
			// 把子树 y 中的 mask 和 val 一个一个地加到 f 中
			for _, p := range singleY {
				update(p.mask, p.val)
			}
		}

		if mask > 0 {
			update(mask, val)
			single = append(single, pair{mask, val})
		}

		mx := 0
		for _, s := range f {
			mx = max(mx, s)
		}
		ans += mx

		return
	}
	dfs(0)
	return ans % mod
}

func goodSubtreeSum(vals, par []int) (ans int) {
	const mod = 1_000_000_007
	const D = 10
	n := len(par)
	g := make([][]int, n)
	for i := 1; i < n; i++ {
		p := par[i]
		g[p] = append(g[p], i)
	}

	type pair struct{ mask, val int }
	var dfs func(int) (map[int]int, []pair)
	dfs = func(x int) (f map[int]int, single []pair) {
		f = map[int]int{}

		// 计算 val 的数位集合 mask
		val := vals[x]
		mask := 0
		for v := val; v > 0; v /= D {
			d := v % D
			if mask>>d&1 > 0 {
				mask = 0
				break
			}
			mask |= 1 << d
		}

		if mask > 0 {
			f[mask] = val
			single = append(single, pair{mask, val})
		}

		for _, y := range g[x] {
			fy, singleY := dfs(y)

			// 启发式合并
			if len(singleY) > len(single) {
				single, singleY = singleY, single
				f, fy = fy, f
			}
			
			single = append(single, singleY...)
			
			// 把子树 y 中的 mask 和 val 一个一个地加到 f 中
			for _, p := range singleY {
				msk, v := p.mask, p.val
				if v <= f[msk] {
					continue
				}
				nf := maps.Clone(f)
				nf[msk] = v
				for msk2, s2 := range f {
					if msk&msk2 == 0 {
						nf[msk|msk2] = max(nf[msk|msk2], v+s2)
					}
				}
				f = nf
			}
		}

		mx := 0
		for _, s := range f {
			mx = max(mx, s)
		}
		ans += mx

		return
	}
	dfs(0)
	return ans % mod
}
