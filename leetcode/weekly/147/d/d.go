package main

// https://space.bilibili.com/206214
func stoneGameII(piles []int) int {
	n := len(piles)
	f := make([][]int, n)
	for i := range f {
		f[i] = make([]int, n+1)
	}
	colQ := make([][]int, n+1) // 每列的滑动窗口最小值（滑动窗口向上滑动）

	for i, s := n-1, 0; i >= 0; i-- {
		s += piles[i]
		diagQ := []int{} // 斜向的滑动窗口最小值（滑动窗口向右下滑动）
		for m := 1; m <= i/2+1; m++ {
			if i+m*2 >= n { // 全拿
				f[i][m] = s
				continue
			}

			// f[i+1][m] 进入列窗口
			q := colQ[m]
			for len(q) > 0 && f[q[len(q)-1]][m] >= f[i+1][m] {
				q = q[:len(q)-1]
			}
			q = append(q, i+1)

			// 队首离开列窗口
			if q[0] > i+m {
				q = q[1:]
			}
			colQ[m] = q

			// f[i+m*2-1][m*2-1] 和 f[i+m*2][m*2] 进入斜向窗口
			for x := m*2 - 1; x <= m*2; x++ {
				for len(diagQ) > 0 {
					j := diagQ[len(diagQ)-1]
					if f[i+j][j] < f[i+x][x] {
						break
					}
					diagQ = diagQ[:len(diagQ)-1]
				}
				diagQ = append(diagQ, x)
			}

			// 队首离开斜向窗口
			if diagQ[0] <= m {
				diagQ = diagQ[1:]
			}

			f[i][m] = s - min(f[q[0]][m], f[i+diagQ[0]][diagQ[0]])
		}
	}

	return f[0][1]
}
