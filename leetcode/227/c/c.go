package main

// github.com/EndlessCheng/codeforces-go
func largestMerge(S, T string) string {
	ans, s, t := make([]byte, 0, len(S)+len(T)), []byte(S), []byte(T)
	for {
		if len(s) == 0 {
			ans = append(ans, t...)
			break
		}
		if len(t) == 0 {
			ans = append(ans, s...)
			break
		}
		i := 0
		for ; i < len(s) && i < len(t) && s[i] == t[i]; i++ {
		}
		if i == len(s) || i == len(t) {
			if len(s) > len(t) {
				ans = append(append(ans, s...), t...)
			} else {
				ans = append(append(ans, t...), s...)
			}
			break
		}
		if s[i] > t[i] {
			ans = append(ans, s[0])
			s = s[1:]
		} else {
			ans = append(ans, t[0])
			t = t[1:]
		}
	}
	return string(ans)
}
