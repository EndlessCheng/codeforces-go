package main

// github.com/EndlessCheng/codeforces-go
func maxValue(n string, x int) string {
	y := byte('0' + x)
	i := 0
	if n[0] != '-' {
		for ; i < len(n) && n[i] >= y; i++ {
		}
	} else {
		for i = 1; i < len(n) && n[i] <= y; i++ {
		}
	}
	return n[:i] + string(y) + n[i:]
}
