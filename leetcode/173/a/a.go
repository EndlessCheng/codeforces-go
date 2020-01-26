package main

func removePalindromeSub(s string) (ans int) {
	n := len(s)
	if n == 0 {
		return
	}
	sRB := make([]byte, n)
	for i := range s {
		sRB[i] = s[n-1-i]
	}
	sR := string(sRB)
	if s == sR {
		return 1
	}
	return 2
}
