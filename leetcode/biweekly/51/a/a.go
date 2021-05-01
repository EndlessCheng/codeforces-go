package main

// github.com/EndlessCheng/codeforces-go
func replaceDigits(s string) string {
	t := []byte(s)
	for i := 1; i < len(t); i += 2 {
		t[i] = t[i-1] + t[i]&15
	}
	return string(t)
}
