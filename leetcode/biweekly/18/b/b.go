package main

// github.com/EndlessCheng/codeforces-go
func breakPalindrome(S string) (ans string) {
	if len(S) == 1 {
		return
	}
	s := []byte(S)
	for i, b := range s[:len(s)/2] {
		if b == 'a' {
			s[len(s)-1-i] = 'b'
		} else {
			s[i] = 'a'
		}
		if t := string(s); ans == "" || t < ans {
			ans = t
		}
		if b == 'a' {
			s[len(s)-1-i] = b
		} else {
			s[i] = b
		}
	}
	return
}
