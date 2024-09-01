package main

// https://space.bilibili.com/206214
func maxScore(grid [][]int) int {
	pos := map[int]int{}
	for i, row := range grid {
		for _, x := range row {
			// 保存所有包含 x 的行号（压缩成二进制数）
			pos[x] |= 1 << i
		}
	}

	f := make([]int, 1<<len(grid))
	for x, posMask := range pos {
		for j := range f {
			for t, lb := posMask, 0; t > 0; t ^= lb {
				lb = t & -t    // lb = 1<<k，其中 k 是行号
				if j&lb == 0 { // 没选过第 k 行的数
					f[j] = max(f[j], f[j|lb]+x)
				}
			}
		}
	}
	return f[0]
}

func maxScore2(grid [][]int) int {
	pos := map[int]int{}
	for i, row := range grid {
		for _, x := range row {
			// 保存所有包含 x 的行号（压缩成二进制数）
			pos[x] |= 1 << i
		}
	}

	allNums := make([]int, 0, len(pos))
	for x := range pos {
		allNums = append(allNums, x)
	}

	f := make([][]int, len(allNums)+1)
	for i := range f {
		f[i] = make([]int, 1<<len(grid))
	}
	for i, x := range allNums {
		for j, v := range f[i] {
			f[i+1][j] = v // 不选 x
			for t, lb := pos[x], 0; t > 0; t ^= lb {
				lb = t & -t    // lb = 1<<k，其中 k 是行号
				if j&lb == 0 { // 没选过第 k 行的数
					f[i+1][j] = max(f[i+1][j], f[i][j|lb]+x) // 选第 k 行的 x
				}
			}
		}
	}
	return f[len(allNums)][0]
}

func maxScore3(grid [][]int) int {
	pos := map[int]int{}
	for i, row := range grid {
		for _, x := range row {
			// 保存所有包含 x 的行号（压缩成二进制数）
			pos[x] |= 1 << i
		}
	}

	allNums := make([]int, 0, len(pos))
	for x := range pos {
		allNums = append(allNums, x)
	}

	n := len(allNums)
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, 1<<len(grid))
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i < 0 {
			return 0
		}
		p := &memo[i][j]
		if *p != -1 {
			return *p
		}
		// 不选 x
		res := dfs(i-1, j)
		// 枚举选第 k 行的 x
		x := allNums[i]
		for t, lb := pos[x], 0; t > 0; t ^= lb {
			lb = t & -t    // lb = 1<<k，其中 k 是行号
			if j&lb == 0 { // 没选过第 k 行的数
				res = max(res, dfs(i-1, j|lb)+x)
			}
		}
		*p = res
		return res
	}
	return dfs(n-1, 0)
}
