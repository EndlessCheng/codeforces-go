package main

import (
	"strconv"
)

// github.com/EndlessCheng/codeforces-go
func countLargestGroup(n int) (ans int) {
	s := strconv.Itoa(n)
	m := len(s)
	memo := make([][]int, m)
	for i := range memo {
		memo[i] = make([]int, m*9+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var dfs func(i, left int, limitHigh bool) int
	dfs = func(i, left int, limitHigh bool) (res int) {
		if i == m {
			if left == 0 {
				return 1
			}
			return
		}

		if !limitHigh {
			p := &memo[i][left]
			if *p != -1 {
				return *p
			}
			defer func() { *p = res }()
		}

		// 当前数位至多填 hi
		hi := 9
		if limitHigh {
			hi = int(s[i] - '0')
		}

		for d := 0; d <= min(hi, left); d++ { // 枚举当前数位填 d
			res += dfs(i+1, left-d, limitHigh && d == hi)
		}
		return
	}

	maxCnt := 0
	for target := 1; target <= m*9; target++ { // 枚举目标数位和
		res := dfs(0, target, true)
		if res > maxCnt {
			maxCnt = res
			ans = 1
		} else if res == maxCnt {
			ans++
		}
	}
	return
}

func countLargestGroup1(n int) (ans int) {
	m := len(strconv.Itoa(n))
	cnt := make([]int, m*9+1) // 数位和 <= 9m
	maxCnt := 0
	for i := 1; i <= n; i++ {
		ds := 0
		for x := i; x > 0; x /= 10 {
			ds += x % 10
		}
		cnt[ds]++
		// 维护 maxCnt 以及 maxCnt 的出现次数
		if cnt[ds] > maxCnt {
			maxCnt = cnt[ds]
			ans = 1
		} else if cnt[ds] == maxCnt {
			ans++
		}
	}
	return
}
