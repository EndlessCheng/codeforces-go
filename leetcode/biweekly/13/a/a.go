package main

import "fmt"

// github.com/EndlessCheng/codeforces-go
func encode(n int) string {
	if n == 0 {
		return ""
	}
	l := 0
	for ; n >= 1<<l; l++ {
		n -= 1 << l
	}
	return fmt.Sprintf("%0*b", l, n)
}
