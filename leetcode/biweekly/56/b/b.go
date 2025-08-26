package main

// github.com/EndlessCheng/codeforces-go
type pair struct{ x, y int }
var dirs = []pair{{0, -1}, {0, 1}, {-1, 0}, {1, 0}} // 左右上下

func nearestExit(maze [][]byte, entrance []int) int {
	m, n := len(maze), len(maze[0])
	sx, sy := entrance[0], entrance[1] // 起点
	maze[sx][sy] = 0 // 访问标记
	q := []pair{{sx, sy}}
	for ans := 1; len(q) > 0; ans++ {
		tmp := q
		q = nil
		for _, p := range tmp {
			// 注意起点不算终点，不能在这里判断 p 是不是终点
			for _, d := range dirs { // 访问相邻的格子
				x, y := p.x+d.x, p.y+d.y
				if 0 <= x && x < m && 0 <= y && y < n && maze[x][y] == '.' { // 之前没有访问过
					if x == 0 || y == 0 || x == m-1 || y == n-1 { // 到达终点
						return ans
					}
					maze[x][y] = 0 // 访问标记
					q = append(q, pair{x, y})
				}
			}
		}
	}
	return -1 // 无法到达终点
}
