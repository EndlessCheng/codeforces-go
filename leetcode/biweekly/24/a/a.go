package main

// github.com/EndlessCheng/codeforces-go
func minStartValue(a []int) (ans int) {
	mi, s := 0, 0
	for _, v := range a {
		s += v
		if s < mi {
			mi = s
		}
	}
	return 1 - mi
}
