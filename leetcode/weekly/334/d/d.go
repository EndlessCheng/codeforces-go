package main

import (
	"container/heap"
	"math"
)

// https://space.bilibili.com/206214
var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func minimumTime(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	if grid[0][1] > 1 && grid[1][0] > 1 { // 无法「等待」
		return -1
	}

	dis := make([][]int, m)
	for i := range dis {
		dis[i] = make([]int, n)
		for j := range dis[i] {
			dis[i][j] = math.MaxInt
		}
	}
	dis[0][0] = 0
	h := &hp{{}}
	for { // 可以等待，就一定可以到达终点
		p := heap.Pop(h).(tuple)
		d, i, j := p.d, p.i, p.j
		if i == m-1 && j == n-1 {
			return d
		}
		for _, q := range dirs { // 枚举周围四个格子
			if x, y := i+q.x, j+q.y; 0 <= x && x < m && 0 <= y && y < n {
				nd := max(d+1, grid[x][y])
				nd += (nd - x - y) % 2 // nd 必须和 x+y 同奇偶
				if dis[x][y] == math.MaxInt {
					dis[x][y] = nd // 更新最短路
					heap.Push(h, tuple{nd, x, y})
				}
			}
		}
	}
}

type tuple struct{ d, i, j int }
type hp []tuple
func (h hp) Len() int              { return len(h) }
func (h hp) Less(i, j int) bool    { return h[i].d < h[j].d }
func (h hp) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{})   { *h = append(*h, v.(tuple)) }
func (h *hp) Pop() (v interface{}) { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
func max(a, b int) int { if a < b { return b }; return a }
