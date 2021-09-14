package main

// github.com/EndlessCheng/codeforces-go
func findCenter(e [][]int) int {
	if e[0][0] == e[1][0] || e[0][0] == e[1][1] {
		return e[0][0]
	}
	return e[0][1]
}
