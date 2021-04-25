package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func maxBuilding(n int, r [][]int) (ans int) {
	sort.Slice(r, func(i, j int) bool { return r[i][0] < r[j][0] })
	r = append([][]int{{1, 0}}, r...)
	if r[len(r)-1][0] != n {
		r = append(r, []int{n, n - 1})
	}
	m := len(r)
	for i := 1; i < m; i++ {
		r[i][1] = min(r[i][1], r[i-1][1]+r[i][0]-r[i-1][0])
	}
	for i := m - 2; i >= 0; i-- {
		r[i][1] = min(r[i][1], r[i+1][1]+r[i+1][0]-r[i][0])
	}
	for i := 0; i < m-1; i++ {
		p, q := r[i], r[i+1]
		ans = max(ans, (q[0]-p[0]+p[1]+q[1])/2)
	}
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
