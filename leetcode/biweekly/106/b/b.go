package main

// https://space.bilibili.com/206214
func longestSemiRepetitiveSubstring(s string) int {
	ans, cnt, left := 1, 0, 0
	for right := 1; right < len(s); right++ {
		if s[right] == s[right-1] {
			cnt++
		}
		for cnt > 1 {
			if s[left] == s[left+1] {
				cnt-- // left 离开窗口后，和 left+1 断开
			}
			left++
		}
		ans = max(ans, right-left+1)
	}
	return ans
}
