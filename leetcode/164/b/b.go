package main

func countServers(grid [][]int) int {
	ans := 0
	r := len(grid)
	for i, gi := range grid {
	outer:
		for j, gij := range gi {
			if gij == 1 {
				for k, gik := range gi {
					if k != j && gik == 1 {
						ans++
						continue outer
					}
				}
				for k := 0; k < r; k++ {
					if k != i && grid[k][j] == 1 {
						ans++
						continue outer
					}
				}
			}
		}
	}
	return ans
}

