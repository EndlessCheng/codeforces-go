package main

import "bytes"

// github.com/EndlessCheng/codeforces-go
func findLexSmallestString(s string, a int, b int) string {
	n := len(s)
	step := gcd(b, n)
	g := gcd(a, 10)
	var ans []byte

	for i := 0; i < n; i += step {
		t := []byte(s[i:] + s[:i])
		modify := func(start int) {
			ch := t[start] - '0' // 最靠前的数字，越小越好
			// ch 可以变成的最小值为 ch%g
			// 例如 ch=5，g=2，那么 ch+2+2+2（模 10）后变成 1，不可能变得更小
			// 从 ch 到 ch%g，需要增加 inc，其中 +10 保证 inc 非负（循环中会 %10 保证结果在 [0,9] 中）
			inc := ch%byte(g) + 10 - ch
			for j := start; j < n; j += 2 {
				t[j] = '0' + (t[j]-'0'+inc)%10
			}
		}
		modify(1) // 累加操作（所有奇数下标）
		if step%2 > 0 { // 轮转一次后，原来的偶数下标变成奇数下标，所以能对偶数下标执行累加操作
			modify(0) // 累加操作（所有偶数下标）
		}
		if ans == nil || bytes.Compare(t, ans) < 0 {
			ans = t
		}
	}

	return string(ans)
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
