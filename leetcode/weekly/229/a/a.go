package main

// github.com/EndlessCheng/codeforces-go
func mergeAlternately(x, y string) (ans string) {
	s := make([]byte, 0, len(x)+len(y))
	for x != "" || y != "" {
		if x != "" {
			s = append(s, x[0])
			x = x[1:]
		}
		if y != "" {
			s = append(s, y[0])
			y = y[1:]
		}
	}
	return string(s)
}
