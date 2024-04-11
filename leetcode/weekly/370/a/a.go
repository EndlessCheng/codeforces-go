package main

// https://space.bilibili.com/206214
func findChampion(grid [][]int) int {
next:
	for i, row := range grid {
		for j, x := range row {
			if j != i && x == 0 {
				continue next
			}
		}
		return i
	}
	panic(-1)
}

func findChampion2(grid [][]int) int {
next:
	for j := range grid[0] {
		for i, row := range grid {
			if i != j && row[j] > 0 { // 有队伍可以击败 j
				continue next
			}
		}
		return j
	}
	panic(-1)
}
