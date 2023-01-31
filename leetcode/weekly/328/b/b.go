package main

// https://space.bilibili.com/206214
func rangeAddQueries(n int, queries [][]int) [][]int {
	// 二维差分模板
	diff := make([][]int, n+2)
	for i := range diff {
		diff[i] = make([]int, n+2)
	}
	update := func(r1, c1, r2, c2, x int) {
		diff[r1+1][c1+1] += x
		diff[r1+1][c2+2] -= x
		diff[r2+2][c1+1] -= x
		diff[r2+2][c2+2] += x
	}
	for _, q := range queries {
		update(q[0], q[1], q[2], q[3], 1)
	}

	// 用二维前缀和复原
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			diff[i][j] += diff[i][j-1] + diff[i-1][j] - diff[i-1][j-1]
		}
	}
	// 保留中间 n*n 的部分，即为答案
	diff = diff[1 : n+1]
	for i, row := range diff {
		diff[i] = row[1 : n+1]
	}
	return diff
}
