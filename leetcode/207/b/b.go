package main

// github.com/EndlessCheng/codeforces-go
func maxUniqueSplit(s string) (ans int) {
	has := map[string]bool{}
	var f func(string)
	f = func(s string) {
		if s == "" {
			if len(has) > ans {
				ans = len(has)
			}
			return
		}
		for i := 1; i <= len(s); i++ {
			if t := s[:i]; !has[t] {
				has[t] = true
				f(s[i:])
				delete(has, t)
			}
		}
	}
	f(s)
	return
}
