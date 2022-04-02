package main

import "strings"

// github.com/EndlessCheng/codeforces-go
func numberOfWays(s string) int64 {
	tot0 := strings.Count(s, "0")
	ans, c0 := 0, 0
	for i, ch := range s {
		if ch == '1' {
			ans += c0 * (tot0 - c0) // 对每个 1，统计左边 0 的个数和右边 0 的个数
		} else {
			c1 := i - c0
			ans += c1 * (len(s) - tot0 - c1) // 对每个 0，统计左边 1 的个数和右边 1 的个数
			c0++
		}
	}
	return int64(ans)
}
