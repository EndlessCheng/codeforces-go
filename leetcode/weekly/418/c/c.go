package main

// https://space.bilibili.com/206214
func constructGridLayout(n int, edges [][]int) [][]int {
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	// 每种度数选一个点
	degToNode := [5]int{-1, -1, -1, -1, -1}
	for x, to := range g {
		degToNode[len(to)] = x
	}

	var row []int
	if degToNode[1] != -1 {
		// 矩阵只有一列
		row = []int{degToNode[1]}
	} else if degToNode[4] == -1 {
		// 矩阵只有两列
		x := degToNode[2]
		for _, y := range g[x] {
			if len(g[y]) == 2 {
				row = []int{x, y}
				break
			}
		}
	} else {
		// 矩形至少有三列
		x := degToNode[2]
		row = []int{x}
		pre := x
		x = g[x][0]
		for len(g[x]) > 2 {
			row = append(row, x)
			for _, y := range g[x] {
				if y != pre && len(g[y]) < 4 {
					pre = x
					x = y
					break
				}
			}
		}
		row = append(row, x)
	}

	k := len(row)
	ans := make([][]int, n/k)
	ans[0] = row
	vis := make([]bool, n)
	for _, x := range row {
		vis[x] = true
	}
	for i := 1; i < len(ans); i++ {
		ans[i] = make([]int, k)
		for j, x := range ans[i-1] {
			for _, y := range g[x] {
				// 上左右的邻居都访问过了，没访问过的邻居只会在 x 下面
				if !vis[y] {
					vis[y] = true
					ans[i][j] = y
					break
				}
			}
		}
	}
	return ans
}
