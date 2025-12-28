package main

import "strconv"

// https://space.bilibili.com/206214
func countBalanced(low, high int64) int64 {
	// 最小的满足要求的数是 11
	if high < 11 {
		return 0
	}

	low = max(low, 11)
	lowS := strconv.FormatInt(low, 10)
	highS := strconv.FormatInt(high, 10)
	n := len(highS)
	diffLH := n - len(lowS)
	memo := make([][][2]int64, n)
	for i := range memo {
		// diff 至少 floor(n/2) * 9，至多 ceil(n/2) * 9，值域大小 n * 9
		memo[i] = make([][2]int64, n*9+1)
	}

	var dfs func(int, int, int, bool, bool) int64
	dfs = func(i, diff, parity int, limitLow, limitHigh bool) (res int64) {
		if i == n {
			if diff != 0 { // 不合法
				return 0
			}
			return 1
		}
		if !limitLow && !limitHigh {
			p := &memo[i][diff+n/2*9][parity] // 保证下标非负
			if *p > 0 {
				return *p - 1
			}
			defer func() { *p = res + 1 }() // 记忆化的时候加一，这样 memo 可以初始化成 0
		}

		lo := 0
		if limitLow && i >= diffLH {
			lo = int(lowS[i-diffLH] - '0')
		}
		hi := 9
		if limitHigh {
			hi = int(highS[i] - '0')
		}

		d := lo
		if limitLow && i < diffLH { // 可以不填任何数
			res = dfs(i+1, diff, parity, true, false) // 上界无约束
			d = 1                                     // 下面填数字，至少从 1 开始填
		}

		for ; d <= hi; d++ {
			// 下一个位置奇偶性翻转
			res += dfs(i+1, diff+(parity*2-1)*d, parity^1,
				limitLow && d == lo, limitHigh && d == hi)
		}
		return
	}
	return dfs(0, 0, 1, true, true)
}
