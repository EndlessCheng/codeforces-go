package main

import (
	"index/suffixarray"
	"math"
	"math/rand"
	"slices"
)

// https://space.bilibili.com/206214
func minimumCost(target string, words []string, costs []int) int {
	n := len(target)

	// 多项式字符串哈希（方便计算子串哈希值）
	// 哈希函数 hash(s) = s[0] * base^(n-1) + s[1] * base^(n-2) + ... + s[n-2] * base + s[n-1]
	const mod = 1_070_777_777
	base := 9e8 - rand.Intn(1e8) // 随机 base，防止 hack（注意 Go1.20 之后的版本，每次随机的数都不一样）
	powBase := make([]int, n+1)  // powBase[i] = base^i
	preHash := make([]int, n+1)  // 前缀哈希值 preHash[i] = hash(s[:i])
	powBase[0] = 1
	for i, b := range target {
		powBase[i+1] = powBase[i] * base % mod
		preHash[i+1] = (preHash[i]*base + int(b)) % mod // 秦九韶算法计算多项式哈希
	}
	// 计算子串 s[l:r] 的哈希值，注意这是左闭右开区间 [l,r) 其中 0<=l<=r<=len(s)
	// 空串的哈希值为 0
	// 计算方法类似前缀和
	subHash := func(l, r int) int {
		return ((preHash[r]-preHash[l]*powBase[r-l])%mod + mod) % mod
	}

	minCost := make([]map[int]int, n+1) // words[i] 的哈希值 -> 最小成本
	lens := map[int]struct{}{}          // 所有 words[i] 的长度集合
	for i, w := range words {
		m := len(w)
		lens[m] = struct{}{}
		// 计算 w 的哈希值
		h := 0
		for _, b := range w {
			h = (h*base + int(b)) % mod
		}
		if minCost[m] == nil {
			minCost[m] = map[int]int{}
		}
		if minCost[m][h] == 0 {
			minCost[m][h] = costs[i]
		} else {
			minCost[m][h] = min(minCost[m][h], costs[i])
		}
	}

	// 有 O(√L) 个不同的长度
	sortedLens := make([]int, 0, len(lens))
	for l := range lens {
		sortedLens = append(sortedLens, l)
	}
	slices.Sort(sortedLens)

	f := make([]int, n+1)
	for i := 1; i <= n; i++ {
		f[i] = math.MaxInt / 2
		for _, sz := range sortedLens {
			if sz > i {
				break
			}
			if cost := minCost[sz][subHash(i-sz, i)]; cost > 0 {
				f[i] = min(f[i], f[i-sz]+cost)
			}
		}
	}
	if f[n] == math.MaxInt/2 {
		return -1
	}
	return f[n]
}

func minimumCostSA(target string, words []string, costs []int) int {
	minCost := map[string]uint16{}
	for i, w := range words {
		c := uint16(costs[i])
		if minCost[w] == 0 {
			minCost[w] = c
		} else {
			minCost[w] = min(minCost[w], c)
		}
	}

	n := len(target)
	type pair struct{ l, cost uint16 }
	from := make([][]pair, n+1)
	sa := suffixarray.New([]byte(target))
	for w, c := range minCost {
		for _, l := range sa.Lookup([]byte(w), -1) {
			r := l + len(w)
			from[r] = append(from[r], pair{uint16(l), c})
		}
	}

	f := make([]int, n+1)
	for i := 1; i <= n; i++ {
		f[i] = math.MaxInt / 2
		for _, p := range from[i] {
			f[i] = min(f[i], f[p.l]+int(p.cost))
		}
	}
	if f[n] == math.MaxInt/2 {
		return -1
	}
	return f[n]
}
