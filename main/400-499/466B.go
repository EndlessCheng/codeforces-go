package main

import (
	. "fmt"
	"io"
	"math"
)

// github.com/EndlessCheng/codeforces-go
func CF466B(in io.Reader, out io.Writer) {
	var n, a, b int64
	Fscan(in, &n, &a, &b)
	if a*b >= 6*n {
		Fprint(out, a*b, "\n", a, b)
		return
	}
	swap := false
	if a > b {
		a, b = b, a
		swap = true
	}
	m := int64(math.Ceil(math.Sqrt(float64(6 * n))))
	minA, minB := int64(2e9), int64(2e9)
	for i := a; i <= m; i++ {
		j := (6*n-1)/i + 1
		if j >= b && i*j < minA*minB {
			minA, minB = i, j
		}
	}
	if swap {
		minA, minB = minB, minA
	}
	Fprint(out, minA*minB, "\n", minA, minB)
}

//func main() { CF466B(os.Stdin, os.Stdout) }
