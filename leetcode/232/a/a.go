package main

// github.com/EndlessCheng/codeforces-go
func areAlmostEqual(S, t string) (ans bool) {
	if S == t {
		return true
	}
	s := []byte(S)
	for i := range s {
		for j := i + 1; j < len(s); j++ {
			s[i], s[j] = s[j], s[i]
			if string(s) == t {
				return true
			}
			s[i], s[j] = s[j], s[i]
		}
	}
	return
}
