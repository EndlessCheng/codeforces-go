package main

import (
	"cmp"
	"math"
	"slices"
	"sort"
)

// https://space.bilibili.com/206214
func subarrayMajority(nums []int, queries [][]int) []int {
	n, m := len(nums), len(queries)

	a := slices.Clone(nums)
	slices.Sort(a)
	a = slices.Compact(a)
	indexToValue := make([]int, n)
	for i, x := range nums {
		indexToValue[i] = sort.SearchInts(a, x)
	}

	cnt := make([]int, len(a)+1)
	maxCnt, minVal := 0, 0
	add := func(i int) {
		v := indexToValue[i]
		cnt[v]++
		c := cnt[v]
		x := nums[i]
		if c > maxCnt {
			maxCnt, minVal = c, x
		} else if c == maxCnt {
			minVal = min(minVal, x)
		}
	}

	ans := make([]int, m)
	blockSize := int(math.Ceil(float64(n) / math.Sqrt(float64(m))))
	type query struct{ bid, l, r, threshold, qid int } // [l,r) 左闭右开
	qs := []query{}
	for i, q := range queries {
		l, r, threshold := q[0], q[1]+1, q[2] // 左闭右开
		// 大区间离线（保证 l 和 r 不在同一个块中）
		if r-l > blockSize {
			qs = append(qs, query{l / blockSize, l, r, threshold, i})
			continue
		}
		// 小区间暴力
		for j := l; j < r; j++ {
			add(j)
		}
		if maxCnt >= threshold {
			ans[i] = minVal
		} else {
			ans[i] = -1
		}
		// 重置数据
		for _, v := range indexToValue[l:r] {
			cnt[v]--
		}
		maxCnt = 0
	}

	slices.SortFunc(qs, func(a, b query) int { return cmp.Or(a.bid-b.bid, a.r-b.r) })

	var r int
	for i, q := range qs {
		l0 := (q.bid + 1) * blockSize
		if i == 0 || q.bid > qs[i-1].bid { // 遍历到一个新的块
			r = l0 // 右端点移动的起点
			// 重置数据
			clear(cnt)
			maxCnt = 0
		}

		// 右端点从 r 移动到 q.r（q.r 不计入）
		for ; r < q.r; r++ {
			add(r)
		}

		// 左端点从 l0 移动到 q.l（l0 不计入）
		tmpMaxCnt, tmpMinVal := maxCnt, minVal
		for l := q.l; l < l0; l++ {
			add(l)
		}
		if maxCnt >= q.threshold {
			ans[q.qid] = minVal
		} else {
			ans[q.qid] = -1
		}

		// 回滚
		maxCnt, minVal = tmpMaxCnt, tmpMinVal
		for _, v := range indexToValue[q.l:l0] {
			cnt[v]--
		}
	}
	return ans
}
