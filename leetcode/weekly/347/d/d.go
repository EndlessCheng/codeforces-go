package main

import "sort"

// https://space.bilibili.com/206214
func maxIncreasingCells(mat [][]int) (ans int) {
	type pair struct{ x, y int }
	g := map[int][]pair{} // 相同元素放在同一组，统计位置
	for i, row := range mat {
		for j, x := range row {
			g[x] = append(g[x], pair{i, j})
		}
	}
	a := make([]int, 0, len(g))
	for k := range g {
		a = append(a, k)
	}
	sort.Ints(a) // 从小到大

	rowMax := make([]int, len(mat))
	colMax := make([]int, len(mat[0]))
	for _, x := range a {
		pos := g[x]
		mx := make([]int, len(pos))
		for i, p := range pos {
			mx[i] = max(rowMax[p.x], colMax[p.y]) + 1 // 先把最大值算出来，再更新 rowMax 和 colMax
			ans = max(ans, mx[i])
		}
		for i, p := range pos {
			rowMax[p.x] = max(rowMax[p.x], mx[i]) // 更新第 p.x 行的最大 f 值
			colMax[p.y] = max(colMax[p.y], mx[i]) // 更新第 p.y 列的最大 f 值
		}
	}
	return
}

func max(a, b int) int { if b > a { return b }; return a }
