package main

// github.com/EndlessCheng/codeforces-go
func distinctEchoSubstrings(s string) (ans int) {
	has := map[string]bool{}
	for i := range s {
		for j := i + 1; 2*j-i <= len(s); j++ {
			// 可以用哈希或后缀数组加速
			t := s[i:j]
			if t == s[j:2*j-i] && !has[t] {
				has[t] = true
			}
		}
	}
	return len(has)
}
