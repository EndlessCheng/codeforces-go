package main

// https://space.bilibili.com/206214
func calc(mat [][]int) int {
	m, n := len(mat), len(mat[0])

	// 221. 最大正方形
	// 计算 mat 下半部分的最大正方形的边长
	sufMax := make([]int, m)
	f := make([]int, n+1)
	mx := 0
	for i := m - 1; i > 0; i-- {
		last := 0
		for j, x := range mat[i] {
			if x == 1 {
				tmp := f[j+1]
				f[j+1] = min(last, f[j+1], f[j]) + 1
				last = tmp
				mx = max(mx, f[j+1])
			} else {
				f[j+1] = 0
				last = 0
			}
		}
		sufMax[i] = mx
	}

	ans := 0
	// 计算 mat 上半部分的最大正方形的边长
	preMax := 0
	clear(f)
	for i, row := range mat[:m-1] {
		last := 0
		for j, x := range row {
			if x == 1 {
				tmp := f[j+1]
				f[j+1] = min(last, f[j+1], f[j]) + 1
				last = tmp
				preMax = max(preMax, f[j+1])
			} else {
				f[j+1] = 0
				last = 0
			}
		}
		if sufMax[i+1] <= ans {
			break // 最优性优化：继续循环不会让 ans 变大
		}
		ans = max(ans, min(preMax, sufMax[i+1])) // 题目要求两个正方形的边长相等
	}

	return ans * ans
}

// 转置矩阵 mat
func transpose(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])
	a := make([][]int, n)
	for i := range a {
		a[i] = make([]int, m)
		for j, row := range mat {
			a[i][j] = row[i]
		}
	}
	return a
}

func maxArea(mat [][]int) int {
	return max(calc(mat), calc(transpose(mat)))
}
