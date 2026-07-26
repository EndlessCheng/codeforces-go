package main

import (
	"container/heap"
	"math"
)

// https://space.bilibili.com/206214
// 奇数下标 1,3 对应向右或向下
// 偶数下标 0,2 对应向左或向上
var dirs = []struct{ x, y int }{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}

func minCost(m, n int, penalty [][]int) int64 {
	dis := make([][][2]int, m)
	for i := range dis {
		dis[i] = make([][2]int, n)
		for j := range dis[i] {
			dis[i][j] = [2]int{math.MaxInt, math.MaxInt}
		}
	}

	// 支付 1 的入口代价
	dis[0][0][1] = 1
	h := &hp{{1, 0, 0, 1}}

	for {
		top := heap.Pop(h).(tuple)
		d, i, j, k := top.dis, top.x, top.y, top.k
		if i == m-1 && j == n-1 {
			return int64(d)
		}
		if d > dis[i][j][k] {
			continue
		}
		p := penalty[i][j]

		// 原地不动
		newDis := d + p
		if newDis < dis[i][j][k^1] {
			dis[i][j][k^1] = newDis
			heap.Push(h, tuple{newDis, i, j, k ^ 1}) // k^1 切换行动编号的奇偶性
		}

		// 移动一步
		for idx, dir := range dirs {
			x, y := i+dir.x, j+dir.y
			if 0 <= x && x < m && 0 <= y && y < n {
				// 如果 k 和 idx 的奇偶性不同，那么违反了奇偶性规则，需要额外支付 p 的代价
				newDis = d + (x+1)*(y+1) + (idx%2^k)*p
				if newDis < dis[x][y][k^1] {
					dis[x][y][k^1] = newDis
					heap.Push(h, tuple{newDis, x, y, k ^ 1}) // k^1 切换行动编号的奇偶性
				}
			}
		}
	}
}

type tuple struct{ dis, x, y, k int }
type hp []tuple

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].dis < h[j].dis }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(tuple)) }
func (h *hp) Pop() (v any)      { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
