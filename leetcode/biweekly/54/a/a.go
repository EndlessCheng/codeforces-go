package main

// github.com/EndlessCheng/codeforces-go
func isCovered(a [][]int, l, r int) bool {
	d := [52]int{}
	for _, p := range a {
		d[p[0]]++
		d[p[1]+1]--
	}
	s := 0
	for _, v := range d[:l] {
		s += v
	}
	for _, v := range d[l : r+1] {
		if s += v; s == 0 {
			return false
		}
	}
	return true
}
