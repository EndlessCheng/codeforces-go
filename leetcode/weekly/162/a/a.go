package main

func oddCells(n int, m int, indices [][]int) (cnt int) {
	x := [50][50]int{}
	for _, ind := range indices {
		row, col := ind[0], ind[1]
		for j := 0; j < m; j++ {
			x[row][j]++
		}
		for i := 0; i < n; i++ {
			x[i][col]++
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if x[i][j]&1 == 1 {
				cnt++
			}
		}
	}
	return
}
