package main

// github.com/EndlessCheng/codeforces-go
func wwork(_, _ int, a []int, b [][]int) (s int64) {
	for _, e := range b {
		z := e[2]
		a[e[0]-1] += z
		a[e[1]-1] += z
		s -= int64(z)
	}
	for _, v := range a {
		if v > 0 {
			s += int64(v)
		}
	}
	return
}
