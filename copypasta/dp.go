package copypasta

/*
若使用滚动数组，注意在下次复用时初始化第一排所有元素
但是实际情况是使用滚动数组仅降低了内存，执行效率与不使用时无异
*/

func dpCollections() {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	_ = min
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	generalDP := func(x, y int) int {
		type pair struct{ x, y int }
		dp := map[pair]int{}
		var f func(x, y int) int
		f = func(x, y int) (ans int) {
			p := pair{x, y}
			if v, ok := dp[p]; ok {
				return v
			}
			defer func() { dp[p] = ans }()
			// ...
			return
		}
		return f(x, y)
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

	_ = []interface{}{generalDP, knapsack01}
}
