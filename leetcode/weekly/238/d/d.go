package main

import "slices"

// https://space.bilibili.com/206214
func maxBuilding(n int, restrictions [][]int) int {
	m := len(restrictions)
	if m == 0 {
		return n - 1
	}

	slices.SortFunc(restrictions, func(a, b []int) int { return a[0] - b[0] })

	// h[i] 表示编号为 id[i] 的建筑的最大高度
	h := make([]int, m)
	h[0] = min(restrictions[0][0]-1, restrictions[0][1])
	for i := 1; i < m; i++ {
		h[i] = min(h[i-1]+restrictions[i][0]-restrictions[i-1][0], restrictions[i][1])
	}
	for i := m - 2; i >= 0; i-- {
		h[i] = min(h[i], h[i+1]+restrictions[i+1][0]-restrictions[i][0])
	}

	ans := max((restrictions[0][0]-1+h[0])/2, h[m-1]+n-restrictions[m-1][0])
	for i := range m - 1 {
		ans = max(ans, (restrictions[i+1][0]-restrictions[i][0]+h[i]+h[i+1])/2)
	}
	return ans
}
