package main

import (
	"strconv"
)

// https://space.bilibili.com/206214
func beautifulNumbers(l, r int) int {
	low := strconv.Itoa(l)
	high := strconv.Itoa(r)
	n := len(high)
	diffLH := n - len(low) // 这样写无需加前导零，也无需 isNum 参数

	type tuple struct{ i, m, s int }
	memo := map[tuple]int{}
	var dfs func(int, int, int, bool, bool) int
	dfs = func(i, m, s int, limitLow, limitHigh bool) (res int) {
		if i == n {
			if s == 0 || m%s > 0 {
				return 0
			}
			return 1
		}
		if !limitLow && !limitHigh {
			t := tuple{i, m, s}
			if v, ok := memo[t]; ok {
				return v
			}
			defer func() { memo[t] = res }()
		}

		lo := 0
		if limitLow && i >= diffLH {
			lo = int(low[i-diffLH] - '0')
		}
		hi := 9
		if limitHigh {
			hi = int(high[i] - '0')
		}

		d := lo
		if limitLow && i < diffLH {
			res = dfs(i+1, 1, 0, true, false) // 什么也不填
			d = 1 // 下面循环从 1 开始
		}
		// 枚举填数位 d
		for ; d <= hi; d++ {
			res += dfs(i+1, m*d, s+d, limitLow && d == lo, limitHigh && d == hi)
		}
		return
	}
	return dfs(0, 1, 0, true, true)
}
