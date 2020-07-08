package main

import (
	. "fmt"
	"math"
)

// github.com/EndlessCheng/codeforces-go
func main() {
	n:=0
	Scan(&n)
	m := int(math.Sqrt(float64(n)) + 1e-8)
	s := -m * (m + 1) * (2*m + 1) / 6
	for i := 1; i <= m; i++ {
		s += i * (i + n/i) * (n/i + 1 - i)
	}
	Print(s)
}
