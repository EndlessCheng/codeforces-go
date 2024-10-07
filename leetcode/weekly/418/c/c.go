package main

// https://space.bilibili.com/206214
func constructGridLayout(n int, edges [][]int) [][]int {
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	// 找一个度数最小的点
	x := 0
	for i, to := range g {
		if len(to) < len(g[x]) {
			x = i
		}
	}

	row := []int{x}
	vis := make([]bool, n)
	vis[x] = true
	degSt := len(g[x]) // 起点的度数
	for { // 注意题目保证 n >= 2，可以至少循环一次
		nxt := -1
		for _, y := range g[x] {
			if !vis[y] && (nxt < 0 || len(g[y]) < len(g[nxt])) {
				nxt = y
			}
		}
		x = nxt
		row = append(row, x)
		vis[x] = true
		if len(g[x]) == degSt {
			break
		}
	}

	k := len(row)
	ans := make([][]int, n/k)
	ans[0] = row
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
