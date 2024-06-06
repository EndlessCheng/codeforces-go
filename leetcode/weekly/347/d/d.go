package main

import "slices"

// https://space.bilibili.com/206214
func maxIncreasingCells(mat [][]int) int {
	type pair struct{ x, y int }
	g := map[int][]pair{}
	for i, row := range mat {
		for j, x := range row {
			g[x] = append(g[x], pair{i, j}) // 相同元素放在同一组，统计位置
		}
	}
	keys := make([]int, 0, len(g))
	for k := range g {
		keys = append(keys, k)
	}
	slices.Sort(keys)

	rowMax := make([]int, len(mat))
	colMax := make([]int, len(mat[0]))
	for _, x := range keys {
		pos := g[x]
		mx := make([]int, len(pos))
		for i, p := range pos {
			mx[i] = max(rowMax[p.x], colMax[p.y]) + 1 // 先把最大值算出来，再更新 rowMax 和 colMax
		}
		for i, p := range pos {
			rowMax[p.x] = max(rowMax[p.x], mx[i]) // 更新第 p.x 行的最大 f 值
			colMax[p.y] = max(colMax[p.y], mx[i]) // 更新第 p.y 列的最大 f 值
		}
	}
	return slices.Max(rowMax)
}
