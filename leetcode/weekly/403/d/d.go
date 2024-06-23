package main

import "math"

// https://space.bilibili.com/206214
func minimumArea(a [][]int) [][]int {
	m, n := len(a), len(a[0])
	// f[i+1][j+1] 表示包含【左上角为 (0,0) 右下角为 (i,j) 的子矩形】中的所有 1 的最小矩形面积
	f := make([][]int, m+1)
	for i := range f {
		f[i] = make([]int, n+1)
	}
	type data struct{ top, left, right int }
	border := make([]data, n)
	for j := range border {
		border[j].top = -1 // 无
	}

	for i, row := range a {
		left, right := -1, 0
		for j, x := range row {
			if x > 0 {
				if left < 0 {
					left = j
				}
				right = j
			}
			preB := border[j]
			if left < 0 { // 这一排目前全是 0
				f[i+1][j+1] = f[i][j+1] // 等于上面的结果
			} else if preB.top < 0 { // 这一排有 1，上面全是 0
				f[i+1][j+1] = right - left + 1
				border[j] = data{i, left, right}
			} else { // 这一排有 1，上面也有 1
				l, r := min(preB.left, left), max(preB.right, right)
				f[i+1][j+1] = (r - l + 1) * (i - preB.top + 1)
				border[j] = data{preB.top, l, r}
			}
		}
	}
	return f
}

func minimumSum(grid [][]int) int {
	ans := math.MaxInt
	f := func(a [][]int) {
		m, n := len(a), len(a[0])
		type pair struct{ l, r int }
		lr := make([]pair, m) // 每一行最左最右 1 的列号
		for i, row := range a {
			l, r := -1, 0
			for j, x := range row {
				if x > 0 {
					if l < 0 {
						l = j
					}
					r = j
				}
			}
			lr[i] = pair{l, r}
		}

		// lt[i+1][j+1] = 包含【左上角为 (0,0) 右下角为 (i,j) 的子矩形】中的所有 1 的最小矩形面积
		lt := minimumArea(a)
		a = rotate(a)
		// lb[i][j+1] = 包含【左下角为 (m-1,0) 右上角为 (i,j) 的子矩形】中的所有 1 的最小矩形面积
		lb := rotate(rotate(rotate(minimumArea(a))))
		a = rotate(a)
		// rb[i][j] = 包含【右下角为 (m-1,n-1) 左上角为 (i,j) 的子矩形】中的所有 1 的最小矩形面积
		rb := rotate(rotate(minimumArea(a)))
		a = rotate(a)
		// rt[i+1][j] = 包含【右上角为 (0,n-1) 左下角为 (i,j) 的子矩形】中的所有 1 的最小矩形面积
		rt := rotate(minimumArea(a))

		if m >= 3 {
			for i := 1; i < m; i++ {
				left, right, top, bottom := n, 0, m, 0
				for j := i + 1; j < m; j++ {
					if p := lr[j-1]; p.l >= 0 {
						left = min(left, p.l)
						right = max(right, p.r)
						top = min(top, j-1)
						bottom = j - 1
					}
					// 图片上左
					area := lt[i][n]                                // minimumArea(a[:i], 0, n)
					area += (right - left + 1) * (bottom - top + 1) // minimumArea(a[i:j], 0, n)
					area += lb[j][n]                                // minimumArea(a[j:], 0, n)
					ans = min(ans, area)
				}
			}
		}

		if m >= 2 && n >= 2 {
			for i := 1; i < m; i++ {
				for j := 1; j < n; j++ {
					// 图片上中
					area := lt[i][n] // minimumArea(a[:i], 0, n)
					area += lb[i][j] // minimumArea(a[i:], 0, j)
					area += rb[i][j] // minimumArea(a[i:], j, n)
					ans = min(ans, area)
					// 图片上右
					area = lt[i][j]  // minimumArea(a[:i], 0, j)
					area += rt[i][j] // minimumArea(a[:i], j, n)
					area += lb[i][n] // minimumArea(a[i:], 0, n)
					ans = min(ans, area)
				}
			}
		}
	}
	f(grid)
	f(rotate(grid))
	return ans
}

// 顺时针旋转矩阵 90°
func rotate(a [][]int) [][]int {
	m, n := len(a), len(a[0])
	b := make([][]int, n)
	for i := range b {
		b[i] = make([]int, m)
	}
	for i, row := range a {
		for j, x := range row {
			b[j][m-1-i] = x
		}
	}
	return b
}
