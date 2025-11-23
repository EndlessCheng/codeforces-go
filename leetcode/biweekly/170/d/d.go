package main

import (
	"cmp"
	"strconv"
)

// https://space.bilibili.com/206214
func totalWaviness(num1, num2 int64) int64 {
	lowS := strconv.FormatInt(num1, 10)
	highS := strconv.FormatInt(num2, 10)
	n := len(highS)
	diffLH := n - len(lowS)
	memo := make([][][3][10]int, n)
	for i := range memo {
		memo[i] = make([][3][10]int, n-1) // 一个数至多包含 n-2 个峰或谷
	}

	var dfs func(int, int, int, int, bool, bool) int
	dfs = func(i, waviness, lastCmp, lastDigit int, limitLow, limitHigh bool) (res int) {
		if i == n {
			return waviness
		}
		if !limitLow && !limitHigh {
			p := &memo[i][waviness][lastCmp+1][lastDigit]
			if *p > 0 {
				return *p - 1
			}
			defer func() { *p = res + 1 }()
		}

		lo := 0
		if limitLow && i >= diffLH {
			lo = int(lowS[i-diffLH] - '0')
		}
		hi := 9
		if limitHigh {
			hi = int(highS[i] - '0')
		}

		for d := lo; d <= hi; d++ {
			w := waviness
			c := 0
			if !limitLow || i > diffLH { // 当前填的数不是最高位
				c = cmp.Compare(d, lastDigit)
			}
			if c != 0 && c == -lastCmp { // 形成了一个峰或谷
				w++
			}
			res += dfs(i+1, w, c, d, limitLow && d == lo, limitHigh && d == hi)
		}
		return
	}
	return int64(dfs(0, 0, 0, 0, true, true))
}
