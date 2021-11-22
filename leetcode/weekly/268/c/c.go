package main

import "sort"

// 统计位置 + 二分位置

// github.com/EndlessCheng/codeforces-go
type RangeFreqQuery struct{ pos [1e4 + 1]sort.IntSlice }

func Constructor(arr []int) (q RangeFreqQuery) {
	for i, value := range arr {
		q.pos[value] = append(q.pos[value], i) // 统计 value 在 arr 中的所有下标位置
	}
	return
}

func (q *RangeFreqQuery) Query(left, right, value int) int {
	p := q.pos[value] // value 在 arr 中的所有下标位置
	return p[p.Search(left):].Search(right + 1) // 在下标位置上二分，求 [left,right] 之间的下标个数，即为 value 的频率
}
