package main

func maxValueOfCoins(piles [][]int, k int) int {
	f := make([]int, k+1)
	sumN := 0
	for _, pile := range piles {
		n := len(pile)
		for i := 1; i < n; i++ {
			pile[i] += pile[i-1] // pile 前缀和
		}
		sumN = min(sumN+n, k) // 优化：j 从前 i 个栈的大小之和开始枚举（不超过 k）
		for j := sumN; j > 0; j-- {
			for w, v := range pile[:min(n, j)] {
				f[j] = max(f[j], f[j-w-1]+v)
			}
		}
	}
	return f[k]
}

func maxValueOfCoins2(piles [][]int, k int) int {
	f := make([][]int, len(piles)+1)
	for i := range f {
		f[i] = make([]int, k+1)
	}
	for i, pile := range piles {
		for j := range k + 1 {
			// 不选这一组中的任何物品
			f[i+1][j] = f[i][j]
			// 枚举选哪个
			v := 0
			for w := range min(j, len(pile)) {
				v += pile[w]
				f[i+1][j] = max(f[i+1][j], f[i][j-w-1]+v)
			}
		}
	}
	return f[len(piles)][k]
}

func maxValueOfCoins1(piles [][]int, k int) int {
	n := len(piles)
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, k+1)
	}
	var dfs func(int, int) int
	dfs = func(i, j int) (res int) {
		if i < 0 {
			return
		}
		p := &memo[i][j]
		if *p != 0 { // 之前计算过
			return *p
		}
		defer func() { *p = res }() // 记忆化

		// 不选这一组中的任何物品
		res = dfs(i-1, j)
		// 枚举选哪个
		v := 0
		for w := range min(j, len(piles[i])) {
			v += piles[i][w]
			// w 从 0 开始，物品体积为 w+1
			res = max(res, dfs(i-1, j-w-1)+v)
		}
		return
	}
	return dfs(n-1, k)
}
