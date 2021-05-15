package main

// github.com/EndlessCheng/codeforces-go
func memLeak(a, b int) []int {
	i := 1
	for ; i <= a || i <= b; i++ {
		if a >= b {
			a -= i
		} else {
			b -= i
		}
	}
	return []int{i, a, b}
}
