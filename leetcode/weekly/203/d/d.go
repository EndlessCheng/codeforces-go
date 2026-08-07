package main

import "math"

// github.com/EndlessCheng/codeforces-go
func stoneGameV1(stoneValue []int) int {
	n := len(stoneValue)
	sum := make([]int, n+1) // stoneValue 的前缀和
	for i, v := range stoneValue {
		sum[i+1] = sum[i] + v
	}

	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, n+1)
	}

	// 返回 Alice 在子数组 [i,j) 中获得的最大分数
	var dfs func(int, int) int
	dfs = func(i, j int) (res int) {
		if j-i == 1 { // 只剩下一块石子，游戏结束
			return 0
		}

		p := &memo[i][j]
		if *p > 0 { // 之前计算过
			return *p
		}

		// 把子数组 [i,j) 分成 [i,k) 和 [k,j)
		for k := i + 1; k < j; k++ {
			sumL, sumR := sum[k]-sum[i], sum[j]-sum[k]
			var score int
			if sumL < sumR { // Bob 丢弃 [k,j)，剩下 [i,k)
				score = dfs(i, k) + sumL
			} else if sumL > sumR { // Bob 丢弃 [i,k)，剩下 [k,j)
				score = dfs(k, j) + sumR
			} else { // sumL = sumR，由 Alice 决定丢弃哪边
				score = max(dfs(i, k), dfs(k, j)) + sumL
			}
			res = max(res, score)
		}

		*p = res // 记忆化
		return
	}

	return dfs(0, n)
}

func stoneGameV2(stoneValue []int) int {
	n := len(stoneValue)
	sum := make([]int, n+1) // stoneValue 的前缀和
	for i, v := range stoneValue {
		sum[i+1] = sum[i] + v
	}

	f := make([][]int, n)
	for i := range f {
		f[i] = make([]int, n+1)
	}

	for i := n - 2; i >= 0; i-- {
		for j := i + 2; j <= n; j++ {
			// 把子数组 [i,j) 分成 [i,k) 和 [k,j)
			for k := i + 1; k < j; k++ {
				sumL, sumR := sum[k]-sum[i], sum[j]-sum[k]
				var score int
				if sumL < sumR { // Bob 丢弃 [k,j)，剩下 [i,k)
					score = f[i][k] + sumL
				} else if sumL > sumR { // Bob 丢弃 [i,k)，剩下 [k,j)
					score = f[k][j] + sumR
				} else { // sumL = sumR，由 Alice 决定丢弃哪边
					score = max(f[i][k], f[k][j]) + sumL
				}
				f[i][j] = max(f[i][j], score)
			}
		}
	}

	return f[0][n]
}

func stoneGameV(stoneValue []int) int {
	n := len(stoneValue)
	sum := make([]int, n+1) // stoneValue 的前缀和
	for i, v := range stoneValue {
		sum[i+1] = sum[i] + v
	}

	f := make([]int, n+1)
	sufMax := make([][]int, n+1)
	for i := range sufMax {
		sufMax[i] = make([]int, n+1)
	}

	for i := n - 1; i >= 0; i-- {
		sufMax[i+1][i+1] = math.MinInt
		sufMax[i][i+1] = -sum[i] // f[i][i+1] - sum[i] = 0 - sum[i] = -sum[i]
		preMax := 0
		k := i + 1
		for j := i + 2; j <= n; j++ {
			for sum[k]-sum[i] <= sum[j]-sum[k] {
				preMax = max(preMax, f[k]+sum[k])
				k++
			}
			// 循环结束后 sum[k] - sum[i] > sum[j] - sum[k]
			q := k
			if sum[k-1]-sum[i] == sum[j]-sum[k-1] {
				q--
			}
			f[j] = max(preMax-sum[i], sufMax[q][j]+sum[j])
			sufMax[i][j] = max(sufMax[i+1][j], f[j]-sum[i])
		}
	}

	return f[n]
}
