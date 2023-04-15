package main

import (
	"math"
	"sort"
)

// https://space.bilibili.com/206214
func minimumVisitedCells(grid [][]int) (mn int) {
	m, n := len(grid), len(grid[0])
	type pair struct{ x, i int }
	colSt := make([][]pair, n) // 每列的单调栈
	for i := m - 1; i >= 0; i-- {
		st := []pair{} // 当前行的单调栈
		for j := n - 1; j >= 0; j-- {
			st2 := colSt[j]
			mn = math.MaxInt
			if i == m-1 && j == n-1 { // 特殊情况：已经是终点
				mn = 0
			} else if g := grid[i][j]; g > 0 {
				// 在单调栈上二分
				k := j + g
				k = sort.Search(len(st), func(i int) bool { return st[i].i <= k })
				if k < len(st) {
					mn = min(mn, st[k].x)
				}
				k = i + g
				k = sort.Search(len(st2), func(i int) bool { return st2[i].i <= k })
				if k < len(st2) {
					mn = min(mn, st2[k].x)
				}
			}

			if mn < math.MaxInt {
				mn++ // 加上 (i,j) 这个格子
				// 插入单调栈
				for len(st) > 0 && mn <= st[len(st)-1].x {
					st = st[:len(st)-1]
				}
				st = append(st, pair{mn, j})
				for len(st2) > 0 && mn <= st2[len(st2)-1].x {
					st2 = st2[:len(st2)-1]
				}
				colSt[j] = append(st2, pair{mn, i})
			}
		}
	}
	// 最后一个算出的 mn 就是 f[0][0]
	if mn == math.MaxInt {
		return -1
	}
	return
}

func min(a, b int) int { if b < a { return b }; return a }
