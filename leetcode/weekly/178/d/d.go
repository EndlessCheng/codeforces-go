package main

func minCost(grid [][]int) (ans int) {
	n, m := len(grid), len(grid[0])
	vis := make([][]bool, n)
	for i := range vis {
		vis[i] = make([]bool, m)
	}

	// 另一种写法是建图套 01BFS 模板（指向的位置连 0 边，另外三个方向连 1 边）
	dir4 := [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	checks := [][2]int{}
	var f func(x, y int)
	f = func(x, y int) {
		if x < 0 || x >= n || y < 0 || y >= m || vis[x][y] {
			return
		}
		vis[x][y] = true
		if x == n-1 && y == m-1 {
			return
		}
		checks = append(checks, [2]int{x, y})
		d := dir4[grid[x][y]-1]
		f(x+d[0], y+d[1])
	}
	f(0, 0)

	for !vis[n-1][m-1] {
		ans++
		tmp := checks
		checks = [][2]int{}
		for _, p := range tmp {
			for _, d := range dir4 {
				f(p[0]+d[0], p[1]+d[1])
			}
		}
	}
	return
}
