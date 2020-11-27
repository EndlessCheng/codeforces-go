package main

// 注：复杂度 O(nm) 的做法：
// 定义 dp[i][j] 表示以 s[i] 和 t[j] 结尾的所有子串对中，恰好只有一个字符不同的子串对的数目

// github.com/EndlessCheng/codeforces-go
func countSubstrings(s string, t string) (ans int) {
	for i := range s {
		for j := range t {
			d := 0
			for k := 0; i+k < len(s) && j+k < len(t); k++ {
				if s[i+k] != t[j+k] {
					d++
				}
				if d > 1 {
					break
				}
				ans += d
			}
		}
	}
	return
}
