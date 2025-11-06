package main

// https://space.bilibili.com/206214
func rangeAddQueries(n int, queries [][]int) [][]int {
	// 二维差分
	diff := make([][]int, n)
	for i := range diff {
		diff[i] = make([]int, n)
	}
	for _, q := range queries {
		r1, c1, r2, c2 := q[0], q[1], q[2], q[3]
		diff[r1][c1]++
		if c2 < n-1 {
			diff[r1][c2+1]--
		}
		if r2 < n-1 {
			diff[r2+1][c1]--
		}
		if r2 < n-1 && c2 < n-1 {
			diff[r2+1][c2+1]++
		}
	}

	// 原地计算 diff 的二维前缀和，作为答案
	ans := make([][]int, n)
	for i := range ans {
		ans[i] = make([]int, n)
		for j := range ans[i] {
			diff[i+1][j+1] += diff[i+1][j] + diff[i][j+1] - diff[i][j]
			ans[i][j] = diff[i+1][j+1]
		}
	}
	return diff
}
