package main

import (
	"container/heap"
	"sort"
)

// https://space.bilibili.com/206214
var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func maxPoints(grid [][]int, queries []int) []int {
	m, n := len(grid), len(grid[0])

	// 查询的下标按照查询值从小到大排序，方便离线
	id := make([]int, len(queries))
	for i := range id {
		id[i] = i
	}
	sort.Slice(id, func(i, j int) bool { return queries[id[i]] < queries[id[j]] })

	ans := make([]int, len(queries))
	h := hp{{grid[0][0], 0, 0}}
	grid[0][0] = 0 // 充当 vis 数组的作用
	cnt := 0
	for _, i := range id {
		q := queries[i]
		for len(h) > 0 && h[0].val < q {
			cnt++
			p := heap.Pop(&h).(tuple)
			for _, d := range dirs { // 枚举周围四个格子
				x, y := p.i+d.x, p.j+d.y
				if 0 <= x && x < m && 0 <= y && y < n && grid[x][y] > 0 {
					heap.Push(&h, tuple{grid[x][y], x, y})
					grid[x][y] = 0 // 充当 vis 数组的作用
				}
			}
		}
		ans[i] = cnt
	}
	return ans
}

type tuple struct{ val, i, j int }
type hp []tuple

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].val < h[j].val }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(tuple)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
