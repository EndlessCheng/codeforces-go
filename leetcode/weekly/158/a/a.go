package main

// github.com/EndlessCheng/codeforces-go
func balancedStringSplit(s string) (ans int) {
	d := 0
	for _, ch := range s {
		if ch == 'L' {
			d++
		} else {
			d--
		}
		if d == 0 {
			ans++
		}
	}
	return
}
