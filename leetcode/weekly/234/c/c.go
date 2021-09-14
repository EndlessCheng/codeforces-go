package main

// github.com/EndlessCheng/codeforces-go
func evaluate(s string, knowledge [][]string) string {
	mp := make(map[string]string, len(knowledge))
	for _, p := range knowledge {
		mp[p[0]] = p[1]
	}
	ans := []byte{}
	l := -1
	for i := range s {
		if b := s[i]; b == '(' {
			l = i
		} else if b == ')' {
			res := mp[s[l+1:i]]
			if res == "" {
				res = "?"
			}
			ans = append(ans, res...)
			l = -1
		} else if l < 0 {
			ans = append(ans, b)
		}
	}
	return string(ans)
}
