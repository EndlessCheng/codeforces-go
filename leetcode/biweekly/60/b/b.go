package main

// github.com/EndlessCheng/codeforces-go
func findFarmland(a [][]int) (ans [][]int) {
	n, m := len(a), len(a[0])
	for i, row := range a {
		for j, v := range row {
			if v == 0 || j > 0 && row[j-1] == 1 || i > 0 && a[i-1][j] == 1 {
				continue
			}
			ii := i
			for ; ii+1 < n && a[ii+1][j] == 1; ii++ {
			}
			jj := j
			for ; jj+1 < m && row[jj+1] == 1; jj++ {
			}
			ans = append(ans, []int{i, j, ii, jj})
		}
	}
	return
}
