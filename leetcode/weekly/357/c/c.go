package main

// https://space.bilibili.com/206214
type pair struct{ x, y int }
var dir4 = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func maximumSafenessFactor(grid [][]int) int {
	n := len(grid)
	dis := make([][]int, n)
	q := []pair{}
	for i, row := range grid {
		dis[i] = make([]int, n)
		for j, x := range row {
			if x > 0 {
				q = append(q, pair{i, j})
			} else {
				dis[i][j] = -1
			}
		}
	}

	groups := [][]pair{q}
	// 多源 BFS
	for len(q) > 0 {
		tmp := q
		q = nil
		for _, p := range tmp {
			for _, d := range dir4 {
				x, y := p.x+d.x, p.y+d.y
				if 0 <= x && x < n && 0 <= y && y < n && dis[x][y] < 0 {
					dis[x][y] = len(groups)
					q = append(q, pair{x, y})
				}
			}
		}
		groups = append(groups, q) // 相同 dis 分组记录
	}

	// 并查集模板
	fa := make([]int, n*n)
	for i := range fa {
		fa[i] = i
	}
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}

	// 从大到小枚举答案
	for ans := len(groups) - 2; ans > 0; ans-- {
		for _, p := range groups[ans] {
			i, j := p.x, p.y
			for _, d := range dir4 {
				x, y := p.x+d.x, p.y+d.y
				if 0 <= x && x < n && 0 <= y && y < n && dis[x][y] >= ans {
					fa[find(x*n+y)] = find(i*n + j)
				}
			}
		}
		if find(0) == find(n*n-1) { // 写这里判断更快些
			return ans
		}
	}
	return 0
}
