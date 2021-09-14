package main

// github.com/EndlessCheng/codeforces-go
func makeEqual(a []string) bool {
	c := [26]int{}
	for _, s := range a {
		for _, b := range s {
			c[b-'a']++
		}
	}
	for _, v := range c {
		if v%len(a) > 0 {
			return false
		}
	}
	return true
}
