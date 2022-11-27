package main

import "runtime/debug"

// https://space.bilibili.com/206214
func countPalindromes(s string) (ans int) {
	const mod int = 1e9 + 7
	n := len(s)
	suf := [10]int{}
	suf2 := [10][10]int{}
	for i := n - 1; i >= 0; i-- {
		d := s[i] - '0'
		for j, c := range suf {
			suf2[d][j] += c
		}
		suf[d]++
	}

	pre := [10]int{}
	pre2 := [10][10]int{}
	for _, d := range s[:n-1] {
		d -= '0'
		suf[d]--
		for j, c := range suf {
			suf2[d][j] -= c // 撤销
		}
		for j, sf := range suf2 {
			for k, c := range sf {
				ans += pre2[j][k] * c // 枚举所有字符组合
			}
		}
		for j, c := range pre {
			pre2[d][j] += c
		}
		pre[d]++
	}
	return ans % mod
}

// https://space.bilibili.com/206214
func init() { debug.SetGCPercent(-1) }