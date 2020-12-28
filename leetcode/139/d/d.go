package main

// github.com/EndlessCheng/codeforces-go
func numSubmatrixSumTarget(a [][]int, target int) (ans int) {
	n, m := len(a), len(a[0])
	sum2d := make([][]int, n+1) // TEMPLATE
	sum2d[0] = make([]int, m+1)
	for i, row := range a {
		sum2d[i+1] = make([]int, m+1)
		for j, v := range row {
			sum2d[i+1][j+1] = sum2d[i+1][j] + sum2d[i][j+1] - sum2d[i][j] + v
		}
	}
	// 左闭右开
	query := func(r1, c1, r2, c2 int) int {
		return sum2d[r2][c2] - sum2d[r2][c1] - sum2d[r1][c2] + sum2d[r1][c1]
	}

	for r1 := range a {
		for r2 := r1 + 1; r2 <= n; r2++ {
			mp := map[int]int{0: 1}
			for c2 := 1; c2 <= m; c2++ {
				s := query(r1, 0, r2, c2)
				ans += mp[s-target]
				mp[s]++
			}
		}
	}
	return
}
