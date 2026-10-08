package main

// github.com/EndlessCheng/codeforces-go
func mergeAlternately(word1, word2 string) string {
	n, m := len(word1), len(word2)
	ans := make([]byte, 0, n+m) // 预分配空间
	for i := range max(n, m) {
		if i < n {
			ans = append(ans, word1[i])
		}
		if i < m {
			ans = append(ans, word2[i])
		}
	}
	return string(ans)
}
