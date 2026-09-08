package main

import (
	"container/heap"
	"math"
)

// https://space.bilibili.com/206214
func minCost(grid [][]int, k int) int {
	const inf = math.MaxInt / 2
	m, n := len(grid), len(grid[0])
	d := make([][]int, m)
	u := make([][]int, m)
	r := make([][]int, m)
	l := make([][]int, m)
	for i := range d {
		d[i] = make([]int, n)
		u[i] = make([]int, n)
		r[i] = make([]int, n)
		l[i] = make([]int, n)
		for j := range d[i] {
			d[i][j] = inf
			u[i][j] = inf
			r[i][j] = inf
			l[i][j] = inf
		}
	}
	d[0][0] = grid[0][0]
	r[0][0] = grid[0][0]
	ans := inf

	// 第一轮循环，把状态从 (0,0) 更新到第一行和第一列
	// 第二轮循环，把状态从第一行向下更新到其余行，从第一列向右更新到其余列，这样就算出了转向 1 次的所有情况
	// ……
	for range k + 1 {
		// 原地转向
		for i := range m {
			for j := range n {
				mn := min(d[i][j], u[i][j], r[i][j], l[i][j])
				d[i][j] = mn
				u[i][j] = mn
				r[i][j] = mn
				l[i][j] = mn
			}
		}

		// 向下
		for i := 1; i < m; i++ {
			for j, x := range grid[i] {
				d[i][j] = min(d[i][j], d[i-1][j]+x)
			}
		}

		// 向上
		for i := m - 2; i >= 0; i-- {
			for j, x := range grid[i] {
				u[i][j] = min(u[i][j], u[i+1][j]+x)
			}
		}

		// 向右
		for i, row := range grid {
			for j := 1; j < n; j++ {
				r[i][j] = min(r[i][j], r[i][j-1]+row[j])
			}
		}

		// 向左
		for i, row := range grid {
			for j := n - 2; j >= 0; j-- {
				l[i][j] = min(l[i][j], l[i][j+1]+row[j])
			}
		}

		ans = min(ans, d[m-1][n-1], r[m-1][n-1])
	}

	if ans < inf {
		return ans
	}
	return -1
}

var dirs = []struct{ x, y int }{{0, -1}, {0, 1}, {-1, 0}, {1, 0}} // 左右上下

func minCostMemo(grid [][]int, k0 int) int {
	m, n := len(grid), len(grid[0])
	memo := make([][][][4]int, k0+1)
	for i := range memo {
		memo[i] = make([][][4]int, m)
		for j := range memo[i] {
			memo[i][j] = make([][4]int, n)
			for p := range memo[i][j] {
				for q := range memo[i][j][p] {
					memo[i][j][p][q] = -1
				}
			}
		}
	}

	var dfs func(int, int, int, int) int
	dfs = func(k, i, j, idx int) int {
		if i == 0 && j == 0 {
			return grid[0][0]
		}

		p := &memo[k][i][j][idx]
		if *p != -1 {
			return *p
		}

		res := math.MaxInt / 2
		for newIdx, dir := range dirs {
			x, y := i+dir.x, j+dir.y
			if 0 <= x && x < m && 0 <= y && y < n {
				newK := k
				if newIdx != idx {
					if k == 0 {
						continue
					}
					newK--
				}
				res = min(res, dfs(newK, x, y, newIdx))
			}
		}
		res += grid[i][j]

		*p = res
		return res
	}

	ans := min(dfs(k0, m-1, n-1, 0), dfs(k0, m-1, n-1, 2))
	if ans < math.MaxInt/2 {
		return ans
	}
	return -1
}

//

func minCost1(grid [][]int, k0 int) (ans int) {
	m, n := len(grid), len(grid[0])
	dis := make([][][][4]int, k0+1)
	for k := range dis {
		dis[k] = make([][][4]int, m)
		for i := range dis[k] {
			dis[k][i] = make([][4]int, n)
			for j := range dis[k][i] {
				for idx := range dis[k][i][j] {
					dis[k][i][j][idx] = math.MaxInt / 2
				}
			}
		}
	}

	// 初始方向可以向右（1）或向下（3）
	v := grid[0][0]
	h := hp{{v, k0, 0, 0, 1}, {v, k0, 0, 0, 3}}
	dis[k0][0][0][1] = v
	dis[k0][0][0][3] = v

	for len(h) > 0 {
		top := heap.Pop(&h).(tuple)
		d, k, i, j, idx := top.dis, top.k, top.i, top.j, top.idx
		if i == m-1 && j == n-1 {
			return d
		}
		if d > dis[k][i][j][idx] {
			continue
		}
		for newIdx, dir := range dirs {
			x, y := i+dir.x, j+dir.y
			if 0 <= x && x < m && 0 <= y && y < n {
				newK := k
				if newIdx != idx {
					if k == 0 {
						continue
					}
					newK--
				}
				newD := d + grid[x][y]
				if newD < dis[newK][x][y][newIdx] {
					dis[newK][x][y][newIdx] = newD
					heap.Push(&h, tuple{newD, newK, x, y, newIdx})
				}
			}
		}
	}
	return -1
}

type tuple struct{ dis, k, i, j, idx int }
type hp []tuple

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].dis < h[j].dis }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(tuple)) }
func (h *hp) Pop() (v any)      { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
