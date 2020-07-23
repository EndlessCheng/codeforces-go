package main

// github.com/EndlessCheng/codeforces-go
func solve(x int) bool {
	for i := 0; i < 249; i++ {
		if i*i%1000 == x {
			return true
		}
	}
	return false
}
