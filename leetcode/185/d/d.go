package main

func numOfArrays(n int, m int, K int) (ans int) {
	dp := make([][][]int, n)
	for i := range dp {
		dp[i] = make([][]int, m+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, K+1)
			for k := range dp[i][j] {
				dp[i][j][k] = -1
			}
		}
	}
	const mod int = 1e9 + 7
	var f func(p, max, k int) int
	f = func(p, max, k int) (res int) {
		if max > m || k > K {
			return
		}
		if p == n {
			if k == K {
				return 1
			}
			return
		}
		dv := &dp[p][max][k]
		if *dv >= 0 {
			return *dv
		}
		defer func() { *dv = res }()
		// 注：这层循环可以用前缀和优化成 O(1)
		for i := 1; i <= m; i++ {
			if i <= max {
				res += f(p+1, max, k)
			} else {
				res += f(p+1, i, k+1)
			}
			res %= mod
		}
		return
	}
	return f(0, 0, 0)
}
