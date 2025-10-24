package main

// github.com/EndlessCheng/codeforces-go
func countOperations(x, y int) (ans int) {
	for y > 0 {
		ans += x / y // x 变成 x%y
		x, y = y, x%y
	}
	return
}
