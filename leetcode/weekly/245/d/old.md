定义 $\textit{dp}(n,\textit{fi},\textit{se})$ 为人数为 $n$，两名选手位置分别在 $\textit{fi}$ 和 $\textit{se}$ 时的最早回合数和最晚回合数。

转移时，枚举下一轮两名选手的位置。

为简化判断，可以在枚举前处理一下两名选手的位置：若两名选手均在中间位置右侧，或者第二名选手比第一名选手更靠近端点，则将第一名选手的位置改为第二名选手关于中间位置对称的位置，第二名选手同理。这样第一名选手必位于中间位置左侧，且第一名选手比第二名选手更靠近端点。

然后枚举第一名选手左侧保留多少个人，以及第一名选手和第二名选手中间保留多少个人，这样就可以得到下一轮两名选手的位置。

```go
func earliestAndLatest(n, firstPlayer, secondPlayer int) []int {
	type pair struct{ min, max int } // 最早回合数和最晚回合数
	dp := make([][][]pair, n+1)
	for i := range dp {
		dp[i] = make([][]pair, n)
		for j := range dp[i] {
			dp[i][j] = make([]pair, n)
		}
	}
	var f func(n, fi, se int) pair
	f = func(n, fi, se int) (ans pair) {
		if fi+se == n-1 { // 发生比拼
			return pair{1, 1}
		}
		if fi >= n-1-fi || fi > n-1-se { // 为简化后续枚举过程，在枚举前处理一下两名选手的位置
			fi, se = n-1-se, n-1-fi
		}
		dv := &dp[n][fi][se]
		if dv.min > 0 {
			return *dv
		}
		defer func() { *dv = ans }()
		ans.min = 1e9
		mid := (n + 1) / 2 // 下一轮人数
		var next pair
		for i := 0; i <= fi; i++ { // 枚举第一名选手左侧保留多少个人
			for j := 0; j < min(se, n-1-se)-fi; j++ { // 枚举第一名选手和第二名选手中间保留多少个人
				if se < mid { // 两人同侧（处理位置后都位于中间位置左侧）
					next = f(mid, i, i+j+1)
				} else { // 两人异侧
					next = f(mid, i, i+j+1+(se*2-n+1)/2)
				}
				ans.min = min(ans.min, next.min)
				ans.max = max(ans.max, next.max)
			}
		}
		// 加上当前回合数
		ans.min++
		ans.max++
		return
	}
	res := f(n, firstPlayer-1, secondPlayer-1) // 下标从 0 开始
	return []int{res.min, res.max}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```
