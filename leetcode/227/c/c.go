package main

// github.com/EndlessCheng/codeforces-go
func largestMerge(s, t string) string {
	ans := make([]byte, 0, len(s)+len(t))
	for {
		if s == "" {
			ans = append(ans, t...)
			break
		}
		if t == "" {
			ans = append(ans, s...)
			break
		}
		if s > t {
			ans = append(ans, s[0])
			s = s[1:]
		} else {
			ans = append(ans, t[0])
			t = t[1:]
		}
	}
	return string(ans)
}
