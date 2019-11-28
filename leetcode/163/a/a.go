package main

func shiftGrid(grid [][]int, k int) [][]int {
	n, m := len(grid), len(grid[0])
	g := make([][]int, n)
	for i := range g {
		g[i] = make([]int, m)
	}
	for i, gi := range grid {
		for j, gij := range gi {
			newPos := i*m + j + k
			ni, nj := newPos/m%n, newPos%m
			g[ni][nj] = gij
		}
	}
	return g
}
