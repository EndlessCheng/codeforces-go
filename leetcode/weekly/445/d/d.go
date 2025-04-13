package main

import (
	"fmt"
	"math/big"
	"strings"
)

// https://space.bilibili.com/206214
func trans(s string, b int) string {
	x := big.Int{}
	fmt.Fscan(strings.NewReader(s), &x)
	return x.Text(b) // 转成 b 进制
}

func countNumbers(l, r string, b int) int {
	const mod = 1_000_000_007
	lowS := trans(l, b)
	highS := trans(r, b)
	n := len(highS)
	diffLH := n - len(lowS)
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, b)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var dfs func(int, int, bool, bool) int
	dfs = func(i, pre int, limitLow, limitHigh bool) (res int) {
		if i == n {
			return 1
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
		hi := b - 1
		if limitHigh {
			hi = int(highS[i] - '0')
		}

		for d := max(lo, pre); d <= hi; d++ {
			res += dfs(i+1, d, limitLow && d == lo, limitHigh && d == hi)
		}
		return res % mod
	}
	return dfs(0, 0, true, true)
}
