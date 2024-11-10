package main

import (
	"bytes"
	"strings"
)

// https://space.bilibili.com/206214
func smallestNumber(s string, t int64) string {
	tmp, cnt := int(t), 0
	for _, p := range []int{2, 3, 5, 7} {
		for tmp%p == 0 {
			tmp /= p
			cnt++
		}
	}
	if tmp > 1 { // t 包含其他质因子
		return "-1"
	}

	// 补前导零（至少一个）
	cnt = max(cnt-len(s)+1, 1)
	s = strings.Repeat("0", cnt) + s

	n := len(s)
	ans := bytes.Repeat([]byte{'0'}, n)
	type pair struct{ i, t int }
	vis := map[pair]bool{}

	var dfs func(int, int, bool) bool
	dfs = func(i, t int, isLimit bool) bool {
		if i == n {
			return t == 1
		}
		if !isLimit {
			p := pair{i, t}
			if vis[p] {
				return false
			}
			vis[p] = true
		}

		if isLimit && i < cnt && dfs(i+1, t, true) { // 填 0（跳过）
			return true
		}

		low := 0
		if isLimit {
			low = int(s[i] - '0')
		}
		for d := max(low, 1); d <= 9; d++ {
			ans[i] = '0' + byte(d) // 直接覆盖，无需恢复现场
			if dfs(i+1, t/gcd(t, d), isLimit && d == low) {
				return true
			}
		}
		return false
	}
	dfs(0, int(t), true)

	i := bytes.LastIndexByte(ans, '0')
	return string(ans[i+1:])
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
