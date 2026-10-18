package main

// https://space.bilibili.com/206214
func lastVisitedCell(grid [][]byte) []int {
	m, n := len(grid), len(grid[0])
	x, y := 0, 0
	for {
		nx, ny := x, y
		switch grid[x][y] {
		case 'U': nx--
		case 'D': nx++
		case 'L': ny--
		default: ny++
		}
		if nx < 0 || nx == m || ny < 0 || ny == n || grid[nx][ny] == 0 {
			return []int{x, y}
		}
		grid[x][y] = 0 // 访问标记
		x, y = nx, ny
	}
}
