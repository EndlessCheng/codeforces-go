package main

import . "fmt"

// github.com/EndlessCheng/codeforces-go
func CF630F() {
	var n int64
	Scan(&n)
	Print((n-6)*(n-5)/2*(n-4)/3*(n-3)/4*(n-2)/5*(n-1)/6*n/7 + (n-5)*(n-4)/2*(n-3)/3*(n-2)/4*(n-1)/5*n/6 + (n-4)*(n-3)/2*(n-2)/3*(n-1)/4*n/5)
}
