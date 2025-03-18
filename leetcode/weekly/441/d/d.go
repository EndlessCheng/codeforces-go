package main

import (
	"slices"
	"strconv"
)

// https://space.bilibili.com/206214
type tuple struct{ i, m, s int }

var memo = map[tuple]int{}
var high []byte

func dfs(i, m, s int, isLimit, isNum bool) (res int) {
	if i < 0 {
		if s == 0 || m%s > 0 {
			return 0
		}
		return 1
	}
	if !isLimit && isNum {
		t := tuple{i, m, s}
		if v, ok := memo[t]; ok {
			return v
		}
		defer func() { memo[t] = res }()
	}

	hi := 9
	if isLimit {
		hi = int(high[i] - '0')
	}

	d := 0
	if !isNum {
		res = dfs(i-1, m, s, false, false) // 什么也不填
		d = 1
	}
	// 枚举填数字 d
	for ; d <= hi; d++ {
		res += dfs(i-1, m*d, s+d, isLimit && d == hi, true)
	}
	return
}

func calc(r int) int {
	high = []byte(strconv.Itoa(r))
	slices.Reverse(high)
	return dfs(len(high)-1, 1, 0, true, false)
}

func beautifulNumbers(l, r int) int {
	return calc(r) - calc(l-1)
}
