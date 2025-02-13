package main

import (
	"strconv"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
var s [100_001][46]int

func init() {
	for i := 1; i < len(s); i++ {
		s[i] = s[i-1]
		sum := 0
		for x := i; x > 0; x /= 10 {
			sum += x % 10
		}
		s[i][sum]++
	}
}

func countBalls2(lowLimit, highLimit int) (ans int) {
	for j := 1; j < len(s[0]); j++ {
		ans = max(ans, s[highLimit][j]-s[lowLimit-1][j])
	}
	return
}

func countBalls1(lowLimit, highLimit int) (ans int) {
	cnt := [46]int{}
	for i := lowLimit; i <= highLimit; i++ {
		s := 0
		for x := i; x > 0; x /= 10 {
			s += x % 10
		}
		cnt[s]++
		ans = max(ans, cnt[s])
	}
	return
}

func countBalls(lowLimit, highLimit int) (ans int) {
	highS := strconv.Itoa(highLimit)
	n := len(highS)
	lowS := strconv.Itoa(lowLimit)
	lowS = strings.Repeat("0", n-len(lowS)) + lowS // 补前导零，和 num2 对齐

	m := int(highS[0]-'0') + (n-1)*9 // 数位和的上界
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, m+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	var dfs func(int, int, bool, bool) int
	dfs = func(i, j int, limitLow, limitHigh bool) (res int) {
		if i == n {
			if j == 0 { // 合法
				return 1
			}
			return
		}

		if !limitLow && !limitHigh {
			p := &memo[i][j]
			if *p >= 0 {
				return *p
			}
			defer func() { *p = res }()
		}

		lo := 0
		if limitLow {
			lo = int(lowS[i] - '0')
		}
		hi := 9
		if limitHigh {
			hi = int(highS[i] - '0')
		}

		for d := lo; d <= min(hi, j); d++ { // 枚举当前数位填 d，但不能超过 j
			res += dfs(i+1, j-d, limitLow && d == lo, limitHigh && d == hi)
		}
		return
	}

	for j := 1; j <= m; j++ {
		ans = max(ans, dfs(0, j, true, true))
	}
	return
}
