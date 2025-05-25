package main

// https://space.bilibili.com/206214
func isConsecutive(x, y byte) bool {
	d := abs(int(x) - int(y))
	return d == 1 || d == 25
}

func lexicographicallySmallestString(s string) (ans string) {
	n := len(s)
	canBeEmpty := make([][]bool, n)
	for i := range canBeEmpty {
		canBeEmpty[i] = make([]bool, n)
	}
	for i := n - 2; i >= 0; i-- {
		canBeEmpty[i+1][i] = true // 空串
		for j := i + 1; j < n; j += 2 {
			// 性质 2
			if isConsecutive(s[i], s[j]) && canBeEmpty[i+1][j-1] {
				canBeEmpty[i][j] = true
				continue
			}
			// 性质 3
			for k := i + 1; k < j-1; k += 2 {
				if canBeEmpty[i][k] && canBeEmpty[k+1][j] {
					canBeEmpty[i][j] = true
					break
				}
			}
		}
	}

	f := make([]string, n+1)
	for i := n - 1; i >= 0; i-- {
		// 包含 s[i]
		res := string(s[i]) + f[i+1]
		// 不包含 s[i]，但 s[i] 不能单独消除，必须和其他字符一起消除
		for j := i + 1; j < n; j += 2 {
			if canBeEmpty[i][j] { // 消除 s[i] 到 s[j]
				res = min(res, f[j+1])
			}
		}
		f[i] = res
	}
	return f[0]
}

func abs(x int) int { if x < 0 { return -x }; return x }

func lexicographicallySmallestString1(s string) string {
	n := len(s)
	memoEmpty := make([][]int8, n)
	for i := range memoEmpty {
		memoEmpty[i] = make([]int8, n)
		for j := range memoEmpty[i] {
			memoEmpty[i][j] = -1
		}
	}

	var canBeEmpty func(int, int) int8
	canBeEmpty = func(i, j int) (res int8) {
		if i > j { // 空串
			return 1
		}
		p := &memoEmpty[i][j]
		if *p != -1 {
			return *p
		}
		defer func() { *p = res }()

		// 性质 2
		if isConsecutive(s[i], s[j]) && canBeEmpty(i+1, j-1) > 0 {
			return 1
		}
		// 性质 3
		for k := i + 1; k < j; k += 2 {
			if canBeEmpty(i, k) > 0 && canBeEmpty(k+1, j) > 0 {
				return 1
			}
		}
		return 0
	}

	memoDfs := make([]string, n)
	for i := range memoDfs {
		memoDfs[i] = "?"
	}
	var dfs func(int) string
	dfs = func(i int) string {
		if i == n {
			return ""
		}
		p := &memoDfs[i]
		if *p != "?" {
			return *p
		}

		// 包含 s[i]
		res := string(s[i]) + dfs(i+1)
		// 不包含 s[i]，但 s[i] 不能单独消除，必须和其他字符一起消除
		for j := i + 1; j < n; j += 2 {
			if canBeEmpty(i, j) > 0 { // 消除 s[i] 到 s[j]
				res = min(res, dfs(j+1))
			}
		}
		*p = res
		return res
	}

	return dfs(0)
}
