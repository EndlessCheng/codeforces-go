package main

// https://space.bilibili.com/206214
func rangeAddQueries(n int, queries [][]int) [][]int {
	m := n

	// 二维差分模板
	diff := make([][]int, n+1)
	for i := range diff {
		diff[i] = make([]int, m+1)
	}
	update := func(r1, c1, r2, c2, x int) {
		r2++
		c2++
		diff[r1][c1] += x
		diff[r1][c2] -= x
		diff[r2][c1] -= x
		diff[r2][c2] += x
	}
	for _, q := range queries {
		update(q[0], q[1], q[2], q[3], 1)
	}

	// 用二维前缀和复原
	ans := make([][]int, n+1)
	ans[0] = make([]int, m+1)
	for i, row := range diff[:n] {
		ans[i+1] = make([]int, m+1)
		for j, x := range row[:m] {
			ans[i+1][j+1] = ans[i+1][j] + ans[i][j+1] - ans[i][j] + x
		}
	}
	ans = ans[1:]
	for i, row := range ans {
		ans[i] = row[1:]
	}
	return ans
}
