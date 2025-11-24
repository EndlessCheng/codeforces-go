package main

import (
	"cmp"
	"strconv"
)

// https://space.bilibili.com/206214
func totalWaviness1(num1, num2 int64) int64 {
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

		isNum := !limitLow || i > diffLH // 前面是否填过数字
		for d := lo; d <= hi; d++ {
			w := waviness
			c := 0
			if isNum { // 当前填的数不是最高位
				c = cmp.Compare(d, lastDigit)
			}
			if c*lastCmp < 0 { // 形成了一个峰或谷
				w++
			}
			res += dfs(i+1, w, c, d, limitLow && d == lo, limitHigh && d == hi)
		}
		return
	}
	return int64(dfs(0, 0, 0, 0, true, true))
}

func totalWaviness(num1, num2 int64) int64 {
	lowS := strconv.FormatInt(num1, 10)
	highS := strconv.FormatInt(num2, 10)
	n := len(highS)
	diffLH := n - len(lowS)
	type pair struct{ wavinessSum, numCnt int }
	memo := make([][3][10]pair, n)

	var dfs func(int, int, int, bool, bool) pair
	dfs = func(i, lastCmp, lastDigit int, limitLow, limitHigh bool) (res pair) {
		if i == n {
			return pair{0, 1} // 本题无特殊约束，能递归到终点的都是合法数字
		}
		if !limitLow && !limitHigh {
			p := &memo[i][lastCmp+1][lastDigit]
			if p.numCnt > 0 {
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

		isNum := !limitLow || i > diffLH // 前面是否填过数字
		for d := lo; d <= hi; d++ {
			c := 0
			if isNum { // 当前填的数不是最高位
				c = cmp.Compare(d, lastDigit)
			}
			sub := dfs(i+1, c, d, limitLow && d == lo, limitHigh && d == hi)
			res.wavinessSum += sub.wavinessSum // 累加子树的波动值
			res.numCnt += sub.numCnt // 累加子树的合法数字个数
			if c*lastCmp < 0 { // 形成了一个峰或谷
				res.wavinessSum += sub.numCnt // 这个峰谷会出现在 sub.numCnt 个数字中
			}
		}
		return
	}
	return int64(dfs(0, 0, 0, true, true).wavinessSum)
}
