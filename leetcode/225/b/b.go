package main

// github.com/EndlessCheng/codeforces-go
func minCharacters(s, t string) (ans int) {
	ans = 1e9

	f := func(s, t string) {
		for i := 'a'; i < 'z'; i++ {
			c := 0
			for _, b := range s {
				if b > i {
					c++
				}
			}
			for _, b := range t {
				if b <= i {
					c++
				}
			}
			ans = min(ans, c)
		}
	}
	f(s, t)
	f(t, s)

	for i := 'a'; i <= 'z'; i++ {
		c := 0
		for _, b := range s {
			if b != i {
				c++
			}
		}
		for _, b := range t {
			if b != i {
				c++
			}
		}
		ans = min(ans, c)
	}
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
