package main

import (
	"math"
	"sort"
)

// https://space.bilibili.com/206214
func smallestTrimmedNumbers(nums []string, queries [][]int) (ans []int) {
	for i, q := range queries {
		q[0] |= i << 32 // 把询问的下标整合到 k 里面，相比 append 这样可以避免扩容
	}
	sort.Slice(queries, func(i, j int) bool { return queries[i][1] < queries[j][1] }) // 按 trim 排序

	m := len(nums[0])
	type pair struct { s string; i int }
	ps := make([]pair, len(nums))
	for i, s := range nums {
		ps[i] = pair{s, i}
	}

	ans = make([]int, len(queries))
	p := 1
	for _, q := range queries {
		for ; p <= q[1]; p++ {
			sort.SliceStable(ps, func(i, j int) bool { return ps[i].s[m-p] < ps[j].s[m-p] }) // 只比较第 m-p 个字符的大小
		}
		ans[q[0]>>32] = ps[q[0]&math.MaxUint32-1].i
	}
	return
}

func smallestTrimmedNumbers2(nums []string, queries [][]int) []int {
	ans := make([]int, len(queries))
	type pair struct {
		s string
		i int
	}
	ps := make([]pair, len(nums))
	for i, q := range queries {
		for j, s := range nums {
			ps[j] = pair{s[len(s)-q[1]:], j}
		}
		// 也可以用稳定排序，但是要慢一些 sort.SliceStable(ps, func(i, j int) bool { return ps[i].s < ps[j].s })
		sort.Slice(ps, func(i, j int) bool { a, b := ps[i], ps[j]; return a.s < b.s || a.s == b.s && a.i < b.i })
		ans[i] = ps[q[0]-1].i
	}
	return ans
}
