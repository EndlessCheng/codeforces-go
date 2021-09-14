package main

// github.com/EndlessCheng/codeforces-go
func checkIfPangram(s string) bool {
	m := 0
	for _, b := range s {
		m |= 1 << (b - 'a')
	}
	return m == 1<<26-1
}
