package main

// github.com/EndlessCheng/codeforces-go
func buddyStrings(s, t string) (ans bool) {
	if len(s) != len(t) {
		return
	}
	var c, ca [26]int
	same, d := false, 0
	for i := range s {
		b, b2 := s[i]-'a', t[i]-'a'
		c[b]++
		c[b2]--
		if ca[b]++; ca[b] > 1 {
			same = true
		}
		if b != b2 {
			d++
		}
	}
	for _, v := range c {
		if v != 0 {
			return
		}
	}
	return d == 2 || d == 0 && same
}
