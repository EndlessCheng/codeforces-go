package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func longestCommonSubpath(_ int, paths [][]int) (ans int) {
	s := []int{}
	r := int(1e9)
	for _, p := range paths {
		r = min(r, len(p))
		s = append(s, 1e9) // 用一个不存在 paths 中的数拼接所有路径
		s = append(s, p...)
	}
	n, m := len(s), len(paths)

	// 标记每个元素属于哪条路径
	ids := make([]int, n)
	id := -1
	for i, v := range s {
		if v == 1e9 {
			id++
			ids[i] = m
		} else {
			ids[i] = id
		}
	}

	// 构建后缀数组和高度数组
	sa := make([]int, n+1)
	rank := make([]int, n+1)
	k := 0
	cmp := func(i, j int) bool {
		if rank[i] != rank[j] {
			return rank[i] < rank[j]
		}
		ri, rj := -1, -1
		if i+k <= n {
			ri = rank[i+k]
		}
		if j+k <= n {
			rj = rank[j+k]
		}
		return ri < rj
	}
	tmp := make([]int, n+1)
	for i := 0; i <= n; i++ {
		sa[i] = i
		rank[i] = -1
		if i < n {
			rank[i] = s[i]
		}
	}
	for k = 1; k <= n; k *= 2 {
		sort.Slice(sa, func(i, j int) bool { return cmp(sa[i], sa[j]) })
		tmp[sa[0]] = 0
		for i := 1; i <= n; i++ {
			tmp[sa[i]] = tmp[sa[i-1]]
			if cmp(sa[i-1], sa[i]) {
				tmp[sa[i]]++
			}
		}
		copy(rank, tmp)
	}
	sa = sa[1:]
	rank = make([]int, n)
	for i := range rank {
		rank[sa[i]] = i
	}
	height := make([]int, n)
	h := 0
	for i, rk := range rank {
		if h > 0 {
			h--
		}
		if rk > 0 {
			for j := sa[rk-1]; i+h < n && j+h < n && s[i+h] == s[j+h]; h++ {
			}
		}
		height[rk] = h
	}

	// 二分求答案
	return sort.Search(r, func(limit int) bool {
		limit++
		vis := make([]int, m)
		for i := 1; i < n; {
			if height[i] < limit {
				i++
				continue
			}
			st := i
			cnt := 0
			for ; i < n && height[i] >= limit; i++ {
				// 检查 sa[i] 和 sa[i-1]
				if i := ids[sa[i]]; i < m && vis[i] != st {
					vis[i] = st
					cnt++
				}
				if i := ids[sa[i-1]]; i < m && vis[i] != st {
					vis[i] = st
					cnt++
				}
			}
			// 连续 m 个属于不同路径的后缀长度均不小于 limit
			if cnt == m {
				return false
			}
		}
		return true
	})
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
