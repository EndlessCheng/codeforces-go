package main

// github.com/EndlessCheng/codeforces-go
func containsPattern(a []int, m int, k int) (ans bool) {
	for l := range a {
	o:
		for r := l + m*k; r <= len(a); r++ {
			b := a[l:r]
			for i, v := range b {
				if v != b[i%m] {
					continue o
				}
			}
			return true
		}
	}
	return
}
