package main

func countNegatives(grid [][]int) (ans int) {
	for _, row := range grid {
		for _, v := range row {
			if v < 0 {
				ans++
			}
		}
	}
	return
}
