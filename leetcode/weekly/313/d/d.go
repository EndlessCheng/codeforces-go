package main

// https://space.bilibili.com/206214
func deleteString(s string) int {
	n := len(s)
	if allEqual(s) {
		return n
	}
	lcp := make([][]int, n+1) // lcp[i][j] 表示 s[i:] 和 s[j:] 的最长公共前缀
	lcp[n] = make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		lcp[i] = make([]int, n+1)
		for j := n - 1; j > i; j-- {
			if s[i] == s[j] {
				lcp[i][j] = lcp[i+1][j+1] + 1
			}
		}
	}
	f := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		for j := 1; i+j*2 <= n; j++ {
			if lcp[i][i+j] >= j { // 说明 s[i:i+j] == s[i+j:i+j*2]
				f[i] = max(f[i], f[i+j])
			}
		}
		f[i]++
	}
	return f[0]
}

func allEqual(s string) bool {
	for i := 1; i < len(s); i++ {
		if s[i] != s[0] {
			return false
		}
	}
	return true
}

func max(a, b int) int { if b > a { return b }; return a }
