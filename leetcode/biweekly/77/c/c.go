package main

// github.com/EndlessCheng/codeforces-go
var dirs = []struct{ x, y int }{{0, -1}, {0, 1}, {-1, 0}, {1, 0}} // 左右上下

func countUnguarded(m int, n int, guards [][]int, walls [][]int) (ans int) {
	guarded := make([][]int8, m)
	for i := range guarded {
		guarded[i] = make([]int8, n)
	}

	// 标记警卫格子、墙格子
	for _, g := range guards {
		guarded[g[0]][g[1]] = -1
	}
	for _, w := range walls {
		guarded[w[0]][w[1]] = -1
	}

	// 遍历警卫
	for _, g := range guards {
		// 遍历视线
		for _, d := range dirs {
			// 视线所及之处，被保卫
			x, y := g[0]+d.x, g[1]+d.y
			for 0 <= x && x < m && 0 <= y && y < n && guarded[x][y] != -1 {
				guarded[x][y] = 1 // 被保卫
				x += d.x
				y += d.y
			}
		}
	}

	// 统计没被保卫的格子数
	for _, row := range guarded {
		for _, x := range row {
			if x == 0 { // 没被保卫
				ans++
			}
		}
	}
	return
}
