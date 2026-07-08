package main

// https://space.bilibili.com/206214
func canMakeSubsequence(s, t string) bool {
	n, m := len(s), len(t)
	// s[suf[i]:] 是 t[i:] 的子序列
	suf := make([]int, m+1)
	suf[m] = n
	j := n
	for i := m - 1; i >= 0; i-- {
		if s[j-1] == t[i] {
			j--
			if j == 0 {
				// s 已是 t 的子序列
				return true
			}
		}
		suf[i] = j
	}

	pre := -1
	for i, ch := range t {
		// 此时 s[:pre+1] 是 t[:i] 的子序列
		if pre+2 == suf[i+1] { // 公式推导见题解
			return true
		}
		if s[pre+1] == byte(ch) {
			pre++
		}
	}
	return false
}

func canMakeSubsequence2(s, t string) bool {
	n, m := len(s), len(t)
	// s[i:] 是 t[suf[i]:] 的子序列（如果 suf[i]=-1 则不是子序列）
	suf := make([]int, n+1)
	suf[n] = m
	j := m
	for i := n - 1; i >= 0; i-- {
		// 上一轮循环 s[i+1] 匹配了 t[j]，减一后继续匹配 s[i]
		j--
		for j >= 0 && t[j] != s[i] {
			j--
		}
		suf[i] = j
	}

	if suf[0] >= 0 {
		// s 已是 t 的子序列
		return true
	}

	pre := -1
	for i, ch := range s {
		// 此时 s[:i] 是 t[:pre+1] 的子序列（如果 pre=m 则不是子序列）
		// 修改 s[i]，那么在 pre 和 suf[i+1] 之间，至少要有一个字母
		if suf[i+1]-pre > 1 {
			return true
		}

		// 上一轮循环 s[i-1] 匹配了 t[pre]，加一后继续匹配 s[i]
		pre++
		for pre < m && t[pre] != byte(ch) {
			pre++
		}
	}
	return false
}

func canMakeSubsequence3(s, t string) bool {
	n := len(s)
	if n > len(t) {
		return false
	}

	j0 := 0 // 在不修改的情况下，s 的前缀 [0, j0-1] 是 t 的当前前缀的子序列
	j1 := 0 // 在改过一次的情况下，s 的前缀 [0, j1-1] 是 t 的当前前缀的子序列
	for _, ch := range t {
		// j1 普通匹配
		if s[j1] == byte(ch) {
			j1++
		}

		// 也可以修改 s[j0] 为 ch，强行匹配
		j1 = max(j1, j0+1)

		// j0 普通匹配
		if s[j0] == byte(ch) {
			j0++
		}

		if j1 == n {
			// s 是 t 的子序列
			return true
		}
	}
	return false
}
