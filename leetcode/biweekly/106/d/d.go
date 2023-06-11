package main

// https://space.bilibili.com/206214
func goodSubsetofBinaryMatrix(grid [][]int) []int {
	idx := map[int]int{}
	for i, row := range grid {
		mask := 0
		for j, x := range row {
			mask |= x << j
		}
		idx[mask] = i
	}
	if i, ok := idx[0]; ok {
		return []int{i}
	}
	for x, i := range idx {
		for y, j := range idx {
			if x&y == 0 {
				if i < j {
					return []int{i, j}
				}
				return []int{j, i}
			}
		}
	}
	return nil
}
