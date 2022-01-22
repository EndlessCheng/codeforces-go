package main

// github.com/EndlessCheng/codeforces-go
func findMiddleIndex(a []int) int {
	tot := 0
	for _, v := range a {
		tot += v
	}
	s := 0
	for i, v := range a {
		if s*2 == tot-v {
			return i
		}
		s += v
	}
	return -1
}
