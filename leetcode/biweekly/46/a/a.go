package main

// github.com/EndlessCheng/codeforces-go
func longestNiceSubstring(s string) (ans string) {
	n := len(s)
	for l := n; l > 0; l-- {
	o:
		for r := l; r <= n; r++ {
			seen := ['z' + 1]bool{}
			t := s[r-l : r]
			for _, b := range t {
				seen[b] = true
			}
			for i := 0; i < 26; i++ {
				if seen['A'+i] != seen['a'+i] {
					continue o
				}
			}
			return t
		}
	}
	return
}
