package main

import (
	"slices"
	"sort"
)

// https://space.bilibili.com/206214
func maxSubstringLength(s string, k int) bool {
	if k == 0 { // 提前返回
		return true
	}

	// 记录每种字母的出现位置
	pos := [26][]int{}
	for i, b := range s {
		b -= 'a'
		pos[b] = append(pos[b], i)
	}

	// 构建有向图
	g := [26][]int{}
	for i, p := range pos {
		if p == nil {
			continue
		}
		l, r := p[0], p[len(p)-1]
		for j, q := range pos {
			if j == i {
				continue
			}
			k := sort.SearchInts(q, l)
			// [l,r] 包含第 j 个小写字母
			if k < len(q) && q[k] <= r {
				g[i] = append(g[i], j)
			}
		}
	}

	// 遍历有向图
	vis := [26]bool{}
	var l, r int
	var dfs func(int)
	dfs = func(x int) {
		vis[x] = true
		p := pos[x]
		l = min(l, p[0]) // 合并区间
		r = max(r, p[len(p)-1])
		for _, y := range g[x] {
			if !vis[y] {
				dfs(y)
			}
		}
	}

	intervals := [][2]int{}
	for i, p := range pos {
		if p == nil {
			continue
		}
		// 如果要包含第 i 个小写字母，最终得到的区间是什么？
		vis = [26]bool{}
		l, r = len(s), 0
		dfs(i)
		// 不能选整个 s，即区间 [0,n-1]
		if l > 0 || r < len(s)-1 {
			intervals = append(intervals, [2]int{l, r})
		}
	}

	return maxNonoverlapIntervals(intervals) >= k
}

// 435. 无重叠区间
// 直接计算最多能选多少个区间
func maxNonoverlapIntervals(intervals [][2]int) (ans int) {
	slices.SortFunc(intervals, func(a, b [2]int) int { return a[1] - b[1] })
	preR := 0
	for _, p := range intervals {
		if p[0] >= preR {
			ans++
			preR = p[1]
		}
	}
	return
}
