package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func minInterval(intervals [][]int, queries []int) []int {
	// 按照区间长度由小到大排序，这样每次回答的时候用的就是长度最小的区间
	sort.Slice(intervals, func(i, j int) bool { a, b := intervals[i], intervals[j]; return a[1]-a[0] < b[1]-b[0] })

	m := len(queries)
	type pair struct{ pos, i int }
	qs := make([]pair, m)
	for i, q := range queries {
		qs[i] = pair{q, i}
	}
	// 离线：按查询位置排序
	sort.Slice(qs, func(i, j int) bool { return qs[i].pos < qs[j].pos })

	// 初始化并查集
	fa := make([]int, m+1)
	for i := range fa {
		fa[i] = i
	}
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}

	ans := make([]int, m)
	for i := range ans {
		ans[i] = -1
	}
	// 对每个区间，回答所有在 [l,r] 范围内的询问
	// 由于每次回答询问之后，都将其指向了下一个询问
	// 所以若 i = find(i) 符合 i < m && qs[i].pos <= r 的条件，则必然是一个在 [l,r] 范围内的还没有回答过的询问
	for _, p := range intervals {
		l, r := p[0], p[1]
		length := r - l + 1
		// 二分找区间左端点 l 在询问中的位置
		i := sort.Search(m, func(i int) bool { return qs[i].pos >= l })
		// 回答所有在 [l,r] 范围内的还没有被回答过的询问
		for i = find(i); i < m && qs[i].pos <= r; i = find(i + 1) {
			ans[qs[i].i] = length
			fa[i] = i + 1
		}
	}
	return ans
}
