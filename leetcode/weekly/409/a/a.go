package main

// https://space.bilibili.com/206214
var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}, {1, 1}, {-1, 1}, {-1, -1}, {1, -1}}

type neighborSum [][2]int

func Constructor(grid [][]int) neighborSum {
	n := len(grid)
	s := make(neighborSum, n*n)
	for i, row := range grid {
		for j, v := range row {
			for k, d := range dirs {
				x, y := i+d.x, j+d.y
				if 0 <= x && x < n && 0 <= y && y < n {
					s[v][k/4] += grid[x][y]
				}
			}
		}
	}
	return s
}

func (s neighborSum) AdjacentSum(value int) int {
	return s[value][0]
}

func (s neighborSum) DiagonalSum(value int) int {
	return s[value][1]
}
