package main

// https://space.bilibili.com/206214
func topoSort(k int, edges [][]int) []int {
	g := make([][]int, k)
	inDeg := make([]int, k)
	for _, e := range edges {
		x, y := e[0]-1, e[1]-1 // 顶点编号从 0 开始，方便计算
		g[x] = append(g[x], y)
		inDeg[y]++
	}
	q := make([]int, 0, k)
	orders := q
	for i, d := range inDeg {
		if d == 0 {
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		x := q[0]
		q = q[1:]
		for _, y := range g[x] {
			if inDeg[y]--; inDeg[y] == 0 {
				q = append(q, y)
			}
		}
	}
	if cap(q) > 0 {
		return nil
	}
	return orders[:k]
}

func buildMatrix(k int, rowConditions, colConditions [][]int) [][]int {
	row := topoSort(k, rowConditions)
	col := topoSort(k, colConditions)
	if row == nil || col == nil {
		return nil
	}
	pos := make([]int, k)
	for i, v := range col {
		pos[v] = i
	}
	ans := make([][]int, k)
	for i, x := range row {
		ans[i] = make([]int, k)
		ans[i][pos[x]] = x + 1
	}
	return ans
}
