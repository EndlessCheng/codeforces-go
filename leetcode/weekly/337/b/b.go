package main

// https://space.bilibili.com/206214
func checkValidGrid(grid [][]int) bool {
	type pair struct{ i, j int }
	pos := make([]pair, len(grid)*len(grid))
	for i, row := range grid {
		for j, x := range row {
			pos[x] = pair{i, j}
		}
	}
	if pos[0] != (pair{}) { // 从左上角出发
		return false
	}
	for i := 1; i < len(pos); i++ {
		dx := abs(pos[i].i - pos[i-1].i)
		dy := abs(pos[i].j - pos[i-1].j) // 移动距离
		if (dx != 2 || dy != 1) && (dx != 1 || dy != 2) { // 不合法
			return false
		}
	}
	return true
}

func abs(x int) int { if x < 0 { return -x }; return x }
