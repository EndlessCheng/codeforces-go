package main

// https://space.bilibili.com/206214
func maximumGap(s, t string) (ans int) {
	n := len(s)
	suf := make([]int, n) // s[i:] 是 t[suf[i]:] 的子序列
	j := len(t)
	for i := n - 1; i > 0; i-- {
		// 上一轮循环 s[i+1] 匹配了 t[j]，j 减一后继续寻找下一个匹配
		j--
		for t[j] != s[i] { // 题目保证 s 是 t 的子序列，下标不会越界
			j--
		}
		suf[i] = j
	}

	pre := -1
	for i, ch := range s[:n-1] {
		// 上一轮循环 s[i-1] 匹配了 t[pre]，pre 加一后继续寻找下一个匹配
		pre++
		for t[pre] != byte(ch) {
			pre++
		}
		// 此时 s[:i+1] 是 t[:pre+1] 的子序列
		// 此时 s[i+1:] 是 t[suf[i+1]:] 的子序列
		ans = max(ans, suf[i+1]-pre)
	}
	return
}
