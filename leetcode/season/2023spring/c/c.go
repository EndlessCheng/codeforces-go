package main

import "sort"

// https://space.bilibili.com/206214
func fieldOfGreatestBlessing(forceField [][]int) (ans int) {
	// 1. 统计所有左下和右上坐标
	var xs, ys []int
	for _, f := range forceField {
		i, j, side := f[0], f[1], f[2]
		xs = append(xs, 2*i-side, 2*i+side)
		ys = append(ys, 2*j-side, 2*j+side)
	}

	// 2. 排序去重
	f := func(a []int) []int {
		sort.Ints(a)
		k := 0
		for _, x := range a[1:] {
			if a[k] != x {
				k++
				a[k] = x
			}
		}
		return a[:k+1]
	}
	xs = f(xs)
	ys = f(ys)

	// 3. 二维差分
	n, m := len(xs), len(ys)
	diff := make([][]int, n+2)
	for i := range diff {
		diff[i] = make([]int, m+2)
	}
	for _, p := range forceField {
		i, j, side := p[0], p[1], p[2]
		r1 := sort.SearchInts(xs, 2*i-side)
		r2 := sort.SearchInts(xs, 2*i+side)
		c1 := sort.SearchInts(ys, 2*j-side)
		c2 := sort.SearchInts(ys, 2*j+side)
		// 将区域 r1<=r<=r2 && c1<=c<=c2 上的数都加上 x
		// 多 +1 是为了方便求后面复原
		diff[r1+1][c1+1]++
		diff[r1+1][c2+2]--
		diff[r2+2][c1+1]--
		diff[r2+2][c2+2]++
	}

	// 4. 直接在 diff 上复原，计算最大值
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			diff[i][j] += diff[i][j-1] + diff[i-1][j] - diff[i-1][j-1]
			ans = max(ans, diff[i][j])
		}
	}
	return
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
