package main

// https://space.bilibili.com/206214
func almostPalindromic(s string) (ans int) {
	n := len(s)
	expand := func(l, r int) {
		for l >= 0 && r < n && s[l] == s[r] {
			l--
			r++
		}
		ans = max(ans, r-l-1)
	}

	for i := range 2*n - 1 {
		l, r := i/2, (i+1)/2
		for l >= 0 && r < n && s[l] == s[r] {
			l--
			r++
		}
		expand(l-1, r) // 删除 s[l]，继续扩展
		expand(l, r+1) // 删除 s[r]，继续扩展
		if ans >= n {  // 优化：提前返回答案
			return n
		}
	}
	return ans
}
