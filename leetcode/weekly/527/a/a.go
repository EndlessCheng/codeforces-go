package main

// https://space.bilibili.com/206214
func findMismatchedSegments(s1, s2 string) (ans [][]string) {
	n := len(s1)
	start := 0
	for i := range n {
		if s1[i] == s2[i] {
			start = i + 1 // 记录段的起点
		} else if i == n-1 || s1[i+1] == s2[i+1] { // i 是这一段的末尾
			ans = append(ans, []string{s1[start : i+1], s2[start : i+1]})
		}
	}
	return
}
