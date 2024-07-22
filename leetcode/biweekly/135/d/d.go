package main

// https://space.bilibili.com/206214
func maximumScore(grid [][]int) (ans int64) {
	n := len(grid)
	colSum := make([][]int64, n)
	for j := range colSum {
		colSum[j] = make([]int64, n+1)
		for i, row := range grid {
			colSum[j][i+1] = colSum[j][i] + int64(row[j])
		}
	}

	f := make([][][2]int64, n)
	for j := range f {
		f[j] = make([][2]int64, n+1)
	}
	for j := 0; j < n-1; j++ {
		// 用前缀最大值优化
		preMax := f[j][0][1] - colSum[j][0]
		for pre := 1; pre <= n; pre++ {
			f[j+1][pre][0] = max(f[j][pre][0], preMax+colSum[j][pre])
			f[j+1][pre][1] = f[j+1][pre][0]
			preMax = max(preMax, f[j][pre][1]-colSum[j][pre])
		}

		// 用后缀最大值优化
		sufMax := f[j][n][0] + colSum[j+1][n]
		for pre := n - 1; pre > 0; pre-- {
			f[j+1][pre][0] = max(f[j+1][pre][0], sufMax-colSum[j+1][pre])
			sufMax = max(sufMax, f[j][pre][0]+colSum[j+1][pre])
		}

		// 单独计算 pre=0 的状态
		f[j+1][0][0] = sufMax // 无需考虑 f[j][0][0]，因为不能连续三列全白
		f[j+1][0][1] = max(f[j][0][0], f[j][n][0]) // 第 j 列要么全白，要么全黑
	}

	for _, row := range f[n-1] {
		ans = max(ans, row[0])
	}
	return ans
}

func maximumScore2(grid [][]int) (ans int64) {
	n := len(grid)
	colSum := make([][]int64, n)
	for j := range colSum {
		colSum[j] = make([]int64, n+1)
		for i, row := range grid {
			colSum[j][i+1] = colSum[j][i] + int64(row[j])
		}
	}

	f := make([][][2]int64, n)
	for j := range f {
		f[j] = make([][2]int64, n+1)
	}
	for j := 0; j < n-1; j++ {
		// pre 表示第 j+1 列的黑格个数
		for pre := 0; pre <= n; pre++ {
			// dec=1 意味着第 j+1 列的黑格个数 (pre) < 第 j+2 列的黑格个数
			for dec := 0; dec < 2; dec++ {
				res := int64(0)
				// 枚举第 j 列有 cur 个黑格
				for cur := 0; cur <= n; cur++ {
					if cur == pre { // 情况一：相等
						// 没有可以计入总分的格子
						res = max(res, f[j][cur][0])
					} else if cur < pre { // 情况二：右边黑格多
						// 第 j 列的第 [cur, pre) 行的格子可以计入总分
						res = max(res, f[j][cur][1]+colSum[j][pre]-colSum[j][cur])
					} else if dec == 0 { // 情况三：cur > pre >= 第 j+2 列的黑格个数
						// 第 j+1 列的第 [pre, cur) 行的格子可以计入总分
						res = max(res, f[j][cur][0]+colSum[j+1][cur]-colSum[j+1][pre])
					} else if pre == 0 { // 情况四（凹形）：cur > pre < 第 j+2 列的黑格个数
						// 此时第 j+2 列全黑最优（递归过程中一定可以枚举到这种情况）
						// 第 j+1 列全白是最优的，所以只需考虑 pre=0 的情况
						// 由于第 j+1 列在 dfs(j+1) 的情况二中已经统计过，这里不重复统计
						res = max(res, f[j][cur][0])
					}
				}
				f[j+1][pre][dec] = res
			}
		}
	}

	for _, row := range f[n-1] {
		ans = max(ans, row[0])
	}
	return ans
}

func maximumScore3(grid [][]int) (ans int64) {
	n := len(grid)
	colSum := make([][]int64, n)
	for j := range colSum {
		colSum[j] = make([]int64, n+1)
		for i, row := range grid {
			colSum[j][i+1] = colSum[j][i] + int64(row[j])
		}
	}

	memo := make([][][2]int64, n-1)
	for i := range memo {
		memo[i] = make([][2]int64, n+1)
		for j := range memo[i] {
			memo[i][j] = [2]int64{-1, -1} // -1 表示没有计算过
		}
	}
	var dfs func(int, int, int) int64
	dfs = func(j, pre, dec int) int64 {
		if j < 0 {
			return 0
		}
		p := &memo[j][pre][dec]
		if *p != -1 { // 之前计算过
			return *p
		}
		res := int64(0)
		// 枚举第 j 列有 cur 个黑格
		for cur := 0; cur <= n; cur++ {
			if cur == pre { // 情况一：相等
				// 没有可以计入总分的格子
				res = max(res, dfs(j-1, cur, 0))
			} else if cur < pre { // 情况二：右边黑格多
				// 第 j 列的第 [cur, pre) 行的格子可以计入总分
				res = max(res, dfs(j-1, cur, 1)+colSum[j][pre]-colSum[j][cur])
			} else if dec == 0 { // 情况三：cur > pre >= 第 j+2 列的黑格个数
				// 第 j+1 列的第 [pre, cur) 行的格子可以计入总分
				res = max(res, dfs(j-1, cur, 0)+colSum[j+1][cur]-colSum[j+1][pre])
			} else if pre == 0 { // 情况四（凹形）：cur > pre < 第 j+2 列的黑格个数
				// 此时第 j+2 列全黑最优（递归过程中一定可以枚举到这种情况）
				// 第 j+1 列全白是最优的，所以只需考虑 pre=0 的情况
				// 由于第 j+1 列在 dfs(j+1) 的情况二中已经统计过，这里不重复统计
				res = max(res, dfs(j-1, cur, 0))
			}
		}
		*p = res // 记忆化
		return res
	}

	// 枚举第 n-1 列有 i 个黑格
	for i := 0; i <= n; i++ {
		ans = max(ans, dfs(n-2, i, 0))
	}
	return
}
