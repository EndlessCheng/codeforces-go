package main

// github.com/EndlessCheng/codeforces-go
func getXORSum(a, b []int) int {
	x := 0
	for _, v := range a {
		x ^= v
	}
	y := 0
	for _, v := range b {
		y ^= v
	}
	return x & y
}
