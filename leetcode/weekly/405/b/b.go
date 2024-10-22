package main

import "fmt"

// https://space.bilibili.com/206214
func validStrings(n int) (ans []string) {
	mask := 1<<n - 1
	for x := range 1 << n {
		if x>>1&x == 0 {
			ans = append(ans, fmt.Sprintf("%0*b", n, mask^x))
		}
	}
	return
}

func validStrings2(n int) (ans []string) {
	path := make([]byte, n)
	var dfs func(i int)
	dfs = func(i int) {
		if i == n {
			ans = append(ans, string(path)) // 注意 string(path) 需要 O(n) 时间
			return
		}

		// 填 1
		path[i] = '1'
		dfs(i + 1)

		// 填 0
		if i == 0 || path[i-1] == '1' {
			path[i] = '0' // 直接覆盖
			dfs(i + 1)
		}
	}
	dfs(0)
	return
}
