package main

import "sort"

func numSubmat(mat [][]int) (ans int) {
	n, m := len(mat), len(mat[0])
	sum := make([][]int, n+1)
	sum[0] = make([]int, m+1)
	for i, row := range mat {
		sum[i+1] = make([]int, m+1)
		for j, v := range row {
			sum[i+1][j+1] = sum[i+1][j] + sum[i][j+1] - sum[i][j] + v
		}
	}
	q := func(r1, c1, r2, c2 int) int { r2++; c2++; return sum[r2][c2] - sum[r2][c1] - sum[r1][c2] + sum[r1][c1] }

	for r2, row := range mat {
		for c2, v := range row {
			if v == 0 {
				continue
			}
			for r1 := 0; r1 <= r2; r1++ {
				c1 := sort.Search(c2+1, func(c1 int) bool { return q(r1, c1, r2, c2) == (r2-r1+1)*(c2-c1+1) })
				ans += c2 - c1 + 1
			}
		}
	}
	return
}
