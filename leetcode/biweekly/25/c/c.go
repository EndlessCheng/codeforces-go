package main

// github.com/EndlessCheng/codeforces-go
func f(s, t string) bool {
	var c1, c2 [26]int
	for _, c := range s {
		c1[c-'a']++
	}
	for _, c := range t {
		c2[c-'a']++
	}
o:
	for i, c := range c1[:] {
		for j := i; j < 26; j++ {
			if c2[j] >= c {
				c2[j] -= c
				continue o
			}
			c -= c2[j]
			c2[j] = 0
		}
		return false
	}
	return true
}

func checkIfCanBreak(s, t string) bool {
	return f(s, t) || f(t, s)
}
