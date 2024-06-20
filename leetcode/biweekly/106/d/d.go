package main

// https://space.bilibili.com/206214
func goodSubsetofBinaryMatrix(grid [][]int) []int {
	n := len(grid[0])
	maskToIdx := make([]int, 1<<n)
	for i := range maskToIdx {
		maskToIdx[i] = -1
	}
	for i, row := range grid {
		mask := 0
		for j, x := range row {
			mask |= x << j
		}
		if mask == 0 {
			return []int{i}
		}
		maskToIdx[mask] = i
	}
	u := 1<<n - 1
	for x, i := range maskToIdx {
		if i < 0 {
			continue
		}
		c := u ^ x
		for y := c; y > 0; y = (y - 1) & c {
			if j := maskToIdx[y]; j >= 0 {
				if i < j {
					return []int{i, j}
				}
				return []int{j, i}
			}
		}
	}
	return nil
}

func goodSubsetofBinaryMatrix2(grid [][]int) []int {
	maskToIdx := map[int]int{}
	for i, row := range grid {
		mask := 0
		for j, x := range row {
			mask |= x << j
		}
		if mask == 0 {
			return []int{i}
		}
		maskToIdx[mask] = i
	}
	for x, i := range maskToIdx {
		for y, j := range maskToIdx {
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
