package main

import (
	"container/heap"
	"math"
	"sort"
)

// https://space.bilibili.com/206214
func minimumVisitedCells(grid [][]int) int {
	colHeaps := make([]hp, len(grid[0])) // 每一列的最小堆
	rowH := hp{}                         // 第 i 行的最小堆
	f := 1                               // 起点算 1 个格子
	for i, row := range grid {
		rowH = rowH[:0]
		for j, g := range row {
			for len(rowH) > 0 && rowH[0].idx < j { // 无法到达第 j 列
				heap.Pop(&rowH) // 弹出无用数据
			}
			colH := colHeaps[j]
			for len(colH) > 0 && colH[0].idx < i { // 无法到达第 i 行
				heap.Pop(&colH) // 弹出无用数据
			}
			if i > 0 || j > 0 {
				f = math.MaxInt
			}
			if len(rowH) > 0 {
				f = rowH[0].f + 1 // 从左边跳过来
			}
			if len(colH) > 0 {
				f = min(f, colH[0].f+1) // 从上边跳过来
			}
			if g > 0 && f < math.MaxInt { // 可以到达 (i, j)
				heap.Push(&rowH, pair{f, g + j}) // 经过的格子数，向右最远能到达的列号
				heap.Push(&colH, pair{f, g + i}) // 经过的格子数，向下最远能到达的行号
			}
			colHeaps[j] = colH
		}
	}
	// 此时的 f 是在 (m-1, n-1) 处算出来的
	if f < math.MaxInt {
		return f
	}
	return -1
}

type pair struct{ f, idx int }
type hp []pair

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].f < h[j].f }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(pair)) }
func (h *hp) Pop() any          { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

func minimumVisitedCells2(grid [][]int) (mn int) {
	m, n := len(grid), len(grid[0])
	type pair struct{ x, i int }
	colStack := make([][]pair, n) // 每列的单调栈
	rowSt := []pair{}             // 行单调栈
	for i := m - 1; i >= 0; i-- {
		rowSt = rowSt[:0]
		for j := n - 1; j >= 0; j-- {
			colSt := colStack[j]
			if i < m-1 || j < n-1 {
				mn = math.MaxInt
			}
			if g := grid[i][j]; g > 0 {
				// 在单调栈上二分
				k := sort.Search(len(rowSt), func(k int) bool { return rowSt[k].i <= j+g })
				if k < len(rowSt) {
					mn = rowSt[k].x
				}
				k = sort.Search(len(colSt), func(k int) bool { return colSt[k].i <= i+g })
				if k < len(colSt) {
					mn = min(mn, colSt[k].x)
				}
			}
			if mn < math.MaxInt {
				mn++ // 加上 (i,j) 这个格子
				// 插入单调栈
				for len(rowSt) > 0 && mn <= rowSt[len(rowSt)-1].x {
					rowSt = rowSt[:len(rowSt)-1]
				}
				rowSt = append(rowSt, pair{mn, j})
				for len(colSt) > 0 && mn <= colSt[len(colSt)-1].x {
					colSt = colSt[:len(colSt)-1]
				}
				colStack[j] = append(colSt, pair{mn, i})
			}
		}
	}
	// 最后一个算出的 mn 就是 f[0][0]
	if mn == math.MaxInt {
		return -1
	}
	return
}
