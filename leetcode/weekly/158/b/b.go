package main

// github.com/EndlessCheng/codeforces-go
var dir8 = []struct{ x, y int }{{1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1}, {0, -1}, {1, -1}}

func queensAttacktheKing(queens [][]int, king []int) (ans [][]int) {
	g := [8][8]bool{}
	for _, q := range queens {
		g[q[0]][q[1]] = true
	}
	for _, d := range dir8 {
		x, y := king[0], king[1]
		for 0 <= x && x < 8 && 0 <= y && y < 8 {
			if g[x][y] {
				ans = append(ans, []int{x, y})
				break
			}
			x += d.x
			y += d.y
		}
	}
	return
}
