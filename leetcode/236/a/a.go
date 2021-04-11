package main

// github.com/EndlessCheng/codeforces-go
func arraySign(a []int) int {
	s := 1
	for _, v := range a {
		if v == 0 {
			return 0
		}
		if v < 0 {
			s *= -1
		}
	}
	return s
}
