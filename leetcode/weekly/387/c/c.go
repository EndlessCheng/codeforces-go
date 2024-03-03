package main

// https://space.bilibili.com/206214
func minimumOperationsToWriteY(grid [][]int) int {
	var cnt1, cnt2 [3]int
	n := len(grid)
	m := n / 2
	for i, row := range grid[:m] {
		cnt1[row[i]]++
		cnt1[row[n-1-i]]++
		for j, x := range row {
			if j != i && j != n-1-i {
				cnt2[x]++
			}
		}
	}
	for _, row := range grid[m:] {
		cnt1[row[m]]++
		for j, x := range row {
			if j != m {
				cnt2[x]++
			}
		}
	}

	maxNotChange := 0
	for i, c1 := range cnt1 {
		for j, c2 := range cnt2 {
			if i != j {
				maxNotChange = max(maxNotChange, c1+c2)
			}
		}
	}
	return n*n - maxNotChange
}
