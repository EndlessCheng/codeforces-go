package main

// github.com/EndlessCheng/codeforces-go
func countUnguarded(m, n int, guards, walls [][]int) (ans int) {
	// 构建网格
	a := make([][]int, m)
	for i := range a {
		a[i] = make([]int, n)
	}
	for _, p := range guards { a[p[0]][p[1]] = 1 }
	for _, p := range walls  { a[p[0]][p[1]] = 2 }

	// 按行模拟
	for _, row := range a {
		for j := 0; j < n; j++ {
			if row[j] == 2 { continue }
			st, has1 := j, false
			// 注意这里的 j 和外层循环的 j 是同一个变量
			for ; j < n && row[j] != 2; j++ {
				if row[j] == 1 {
					has1 = true
				}
			}
			if has1 {
				for ; st < j; st++ {
					if row[st] == 0 {
						row[st] = -1
					}
				}
			}
		}
	}

	// 按列模拟
	for j := 0; j < n; j++ {
		for i := 0; i < m; i++ {
			if a[i][j] == 2 { continue }
			st, has1 := i, false
			// 注意这里的 i 和外层循环的 i 是同一个变量
			for ; i < m && a[i][j] != 2; i++ {
				if a[i][j] == 1 {
					has1 = true
				}
			}
			if has1 {
				for ; st < i; st++ {
					if a[st][j] == 0 {
						a[st][j] = -1
					}
				}
			}
		}
	}

	// 统计答案
	for _, row := range a {
		for _, v := range row {
			if v == 0 {
				ans++
			}
		}
	}
	return
}
