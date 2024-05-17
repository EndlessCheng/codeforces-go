package main

// github.com/EndlessCheng/codeforces-go
func numSubmatrixSumTarget(a [][]int, target int) (ans int) {
	n, m := len(a), len(a[0])
	sum := make([][]int, n+1)
	sum[0] = make([]int, m+1)
	for i, row := range a {
		sum[i+1] = make([]int, m+1)
		for j, v := range row {
			sum[i+1][j+1] = sum[i+1][j] + sum[i][j+1] - sum[i][j] + v
		}
	}
	// 左闭右开
	query := func(r1, c1, r2, c2 int) int {
		return sum[r2][c2] - sum[r2][c1] - sum[r1][c2] + sum[r1][c1]
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
