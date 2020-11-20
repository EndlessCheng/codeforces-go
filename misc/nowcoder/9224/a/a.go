package main

// github.com/EndlessCheng/codeforces-go
func stick(A int64) int {
	f := [100]int{1: 1}
	for i := 2; i < 100; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	i := 1
	for a := int(A); a >= f[i]; i++ {
		a -= f[i]
	}
	return i - 1
}
