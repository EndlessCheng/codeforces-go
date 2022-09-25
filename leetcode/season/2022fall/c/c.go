package main

// https://space.bilibili.com/206214
var dirs = []struct{ x, y int }{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} // 右下左上（顺时针）

func ballGame(num int, plate []string) (ans [][]int) {
	m, n := len(plate), len(plate[0])
	check := func(x, y, d int) bool {
		for left := num; plate[x][y] != 'O'; left-- {
			if left == 0 { // 无剩余步数
				return false
			}
			if plate[x][y] == 'W' { // 逆时针
				d = (d + 3) % 4
			} else if plate[x][y] == 'E' { // 顺时针
				d = (d + 1) % 4
			}
			x += dirs[d].x
			y += dirs[d].y
			if x < 0 || x >= m || y < 0 || y >= n { // 从另一边出去了
				return false
			}
		}
		return true
	}
	for j := 1; j < n-1; j++ {
		if plate[0][j] == '.' && check(0, j, 1) {
			ans = append(ans, []int{0, j})
		}
		if plate[m-1][j] == '.' && check(m-1, j, 3) {
			ans = append(ans, []int{m - 1, j})
		}
	}
	for i := 1; i < m-1; i++ {
		if plate[i][0] == '.' && check(i, 0, 0) {
			ans = append(ans, []int{i, 0})
		}
		if plate[i][n-1] == '.' && check(i, n-1, 2) {
			ans = append(ans, []int{i, n - 1})
		}
	}
	return
}
