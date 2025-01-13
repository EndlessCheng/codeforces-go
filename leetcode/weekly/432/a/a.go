package main

import "slices"

// https://space.bilibili.com/206214
func zigzagTraversal(grid [][]int) (ans []int) {
	n := len(grid[0])
	end := n - 1 - n%2
	for i, row := range grid {
		if i%2 > 0 {
			for j := end; j >= 0; j -= 2 {
				ans = append(ans, row[j])
			}
		} else {
			for j := 0; j < n; j += 2 {
				ans = append(ans, row[j])
			}
		}
	}
	return
}

func zigzagTraversal1(grid [][]int) (ans []int) {
	ok := true
	for i, row := range grid {
		if i%2 > 0 {
			slices.Reverse(row)
		}
		for _, x := range row {
			if ok {
				ans = append(ans, x)
			}
			ok = !ok
		}
	}
	return
}
