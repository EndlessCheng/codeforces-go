package main

// https://space.bilibili.com/206214
func longestSemiRepetitiveSubstring(s string) int {
	ans, left, same := 1, 0, 0
	for right := 1; right < len(s); right++ {
		if s[right] == s[right-1] {
			same++
			if same > 1 {
				left++
				for s[left] != s[left-1] {
					left++
				}
				same = 1
			}
		}
		ans = max(ans, right-left+1)
	}
	return ans
}

func max(a, b int) int { if b > a { return b }; return a }
