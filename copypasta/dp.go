package copypasta

/*
若使用滚动数组，注意在下次复用时初始化第一排所有元素
但是实际情况是使用滚动数组仅降低了内存，执行效率与不使用时无异
*/

func dpCollections() {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	knapsack01 := func(values, weights []int, maxW int) int {
		n := len(values)
		dp := make([][]int, n+1)
		for i := range dp {
			dp[i] = make([]int, maxW+1)
		}
		for i, vi := range values {
			wi := weights[i]
			for j, dpij := range dp[i] {
				if j < wi {
					dp[i+1][j] = dpij
				} else {
					dp[i+1][j] = max(dpij, dp[i][j-wi]+vi)
				}
			}
		}
		return dp[n][maxW]
	}

	_ = []interface{}{knapsack01}
}
