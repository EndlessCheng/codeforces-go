package main

import "strconv"

// https://space.bilibili.com/206214
// 判断数位和 s 是否为好数
func isGood(s int) bool {
	if s < 100 { // s 是个位数或者两位数
		return s/10 != s%10 // 十位和个位不相等即为好数
	}
	// s 是三位数，其百位一定是 1
	return 1 < s/10%10 && s/10%10 < s%10 // 只能严格递增
}

func countFancy(l, r int64) int64 {
	lowS := strconv.FormatInt(l, 10)
	highS := strconv.FormatInt(r, 10)
	n := len(highS)
	diffLH := n - len(lowS)
	memo := make([][][10][4]int64, n)
	for i := range memo {
		memo[i] = make([][10][4]int64, n*9+1) // 数位和最大 9n
	}

	const (
		stateInit    = iota // 已经填了至多一个数（不含前导零）
		stateInc            // 已填数字是严格递增的
		stateDec            // 已填数字是严格递减的
		stateNotGood        // 已填数字不是好数
	)

	var dfs func(int, int, int, int, bool, bool) int64
	dfs = func(i, digitSum, prev, state int, limitLow, limitHigh bool) (res int64) {
		if i == n {
			if state != stateNotGood || isGood(digitSum) {
				return 1 // 合法
			}
			return 0 // 不合法
		}
		if !limitLow && !limitHigh {
			dv := &memo[i][digitSum][prev][state]
			if *dv > 0 {
				return *dv - 1
			}
			defer func() { *dv = res + 1 }() // 这样写无需初始化 memo 为 -1
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
		// 通过 limitLow 和 i 可以判断能否不填数字，无需 isNum 参数
		if limitLow && i < diffLH {
			// 不填数字，上界不受约束
			res = dfs(i+1, 0, 0, stateInit, true, false)
			d = 1 // 下面填数字，从 1 开始填
		}

		for ; d <= hi; d++ {
			newState := state
			switch state {
			case stateInit:
				if prev > 0 { // 之前填过数
					if d > prev {
						newState = stateInc
					} else if d < prev {
						newState = stateDec
					} else {
						newState = stateNotGood
					}
				}
			case stateInc:
				if d <= prev {
					newState = stateNotGood
				}
			case stateDec:
				if d >= prev {
					newState = stateNotGood
				}
			}
			res += dfs(i+1, digitSum+d, d, newState, limitLow && d == lo, limitHigh && d == hi)
		}
		return
	}

	return dfs(0, 0, 0, stateInit, true, true)
}
