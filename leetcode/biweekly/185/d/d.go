package main

import "strconv"

// https://space.bilibili.com/206214
func goodIntegers(l, r int64, k int) int64 {
	lowS := strconv.FormatInt(l, 10)
	highS := strconv.FormatInt(r, 10)
	n := len(highS)
	diffLH := n - len(lowS)
	memo := make([][10]int64, n)
	for i := range memo {
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var dfs func(int, int, bool, bool) int64
	dfs = func(i, pre int, limitLow, limitHigh bool) (res int64) {
		if i == n {
			return 1 // 找到一个好数
		}
		if !limitLow && !limitHigh {
			p := &memo[i][pre]
			if *p >= 0 {
				return *p
			}
			defer func() { *p = res }()
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
		if limitLow && i < diffLH {
			// 不填数字，上界不受约束
			res = dfs(i+1, 0, true, false)
			d = 1 // 下面填数字，从 1 开始填
		}

		// 如果在 diffLH 之前填过数字，那么 limitLow 一定是 false
		isFirst := limitLow && i <= diffLH
		for ; d <= hi; d++ {
			if isFirst || abs(d-pre) <= k {
				res += dfs(i+1, d, limitLow && d == lo, limitHigh && d == hi)
			}
		}
		return
	}

	// pre 的初始值随意
	return dfs(0, 0, true, true)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
