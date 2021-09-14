package main

// github.com/EndlessCheng/codeforces-go
func makeGood(s string) (ans string) {
o:
	for len(s) > 1 {
		for i := 1; i < len(s); i++ {
			if s[i-1] != s[i] && (s[i-1]^s[i])&31 == 0 {
				s = s[:i-1] + s[i+1:]
				continue o
			}
		}
		break
	}
	return s
}
