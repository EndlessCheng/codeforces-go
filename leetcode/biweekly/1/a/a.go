package main

// github.com/EndlessCheng/codeforces-go
func fixedPoint(a []int) (ans int) {
	for i, v := range a {
		if i == v {
			return i
		}
	}
	return -1
}
