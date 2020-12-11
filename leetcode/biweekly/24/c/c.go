package main

// github.com/EndlessCheng/codeforces-go
func getHappyString(n, k int) (ans string) {
	s := make([]byte, n)
	var f func(int)
	f = func(p int) {
		if p == n {
			k--
			return
		}
		for s[p] = 'a'; s[p] < 'd'; s[p]++ {
			if p == 0 || s[p] != s[p-1] {
				f(p + 1)
				if k == 0 {
					return
				}
			}
		}
	}
	f(0)
	if k > 0 {
		return
	}
	return string(s)
}
