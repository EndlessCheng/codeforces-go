package main

func countSquares0(matrix [][]int) (ans int) {
	sum := make([][]int, len(matrix)+1)
	sum[0] = make([]int, len(matrix[0])+1)
	for i, mi := range matrix {
		sum[i+1] = make([]int, len(mi)+1)
		for j, mij := range mi {
			sum[i+1][j+1] = sum[i+1][j] + sum[i][j+1] - sum[i][j] + mij
		}
	}
	query := func(r1, c1, r2, c2 int) int {
		return sum[r2][c2] - sum[r2][c1] - sum[r1][c2] + sum[r1][c1]
	}
	for i, mi := range matrix {
		for j := range mi {
			for sz := 1; i+sz <= len(matrix) && j+sz <= len(mi) && query(i, j, i+sz, j+sz) == sz*sz; sz++ {
				ans++
			}
		}
	}
	return ans
}

func countSquares1(matrix [][]int) (ans int) {
	m, n := len(matrix), len(matrix[0])
	for top := range m { // 枚举上边界
		a := make([]int, n)
		for bottom := top; bottom < m; bottom++ { // 枚举下边界
			h := bottom - top + 1 // 高
			// 2348. 全 h 子数组的数目
			last := -1
			for j := range n {
				a[j] += matrix[bottom][j] // 把 bottom 这一行的值加到 a 中
				if a[j] != h {
					last = j // 记录上一个非 h 元素的位置
				} else if j-last >= h { // 右端点为 j 的长为 h 的子数组全是 h
					ans++
				}
			}
		}
	}
	return
}

func countSquares2(matrix [][]int) (ans int) {
	for i, row := range matrix {
		for j, x := range row {
			if x > 0 && i > 0 && j > 0 {
				row[j] += min(matrix[i-1][j], matrix[i-1][j-1], row[j-1])
			}
			ans += row[j]
		}
	}
	return
}

func countSquare(heights []int) (ans int) {
	st := []int{-1} // 在栈中只有一个数的时候，栈顶的「下面那个数」是 -1，对应 left[i] = -1 的情况
	for r, hr := range heights {
		for len(st) > 1 && heights[st[len(st)-1]] >= hr {
			h := heights[st[len(st)-1]] // 矩形的高
			st = st[:len(st)-1]
			l := st[len(st)-1] // 栈顶下面那个数就是 l
			w := r - l - 1
			upper := min(h, w)
			lower := hr + 1
			if l >= 0 {
				lower = max(heights[l], hr) + 1
			}
			if lower <= upper {
				ans += (w*2 + 2 - lower - upper) * (upper - lower + 1) / 2
			}
		}
		st = append(st, r)
	}
	return
}

func countSquares(matrix [][]int) (ans int) {
	heights := make([]int, len(matrix[0])+1) // 末尾多一个 0，理由见我 84 题题解
	for _, row := range matrix {
		// 计算底边为 row 的柱子高度
		for j, x := range row {
			if x == 0 {
				heights[j] = 0 // 柱子高度为 0
			} else {
				heights[j]++ // 柱子高度加一
			}
		}
		ans += countSquare(heights)
	}
	return
}
