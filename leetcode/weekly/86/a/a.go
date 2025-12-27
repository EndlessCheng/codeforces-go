package main

// https://space.bilibili.com/206214
func numMagicSquaresInside(grid [][]int) (ans int) {
	m, n := len(grid), len(grid[0])
	for i := range m - 2 {
		for j := range n - 2 {
			if grid[i+1][j+1] != 5 {
				continue
			}

			mask := 0
			var rSum, cSum [3]int
			for r, row := range grid[i : i+3] {
				for c, x := range row[j : j+3] {
					mask |= 1 << x
					rSum[r] += x
					cSum[c] += x
				}
			}

			if mask == 1<<10-2 &&
				rSum[0] == 15 && rSum[1] == 15 &&
				cSum[0] == 15 && cSum[1] == 15 {
				ans++
			}
		}
	}
	return
}
