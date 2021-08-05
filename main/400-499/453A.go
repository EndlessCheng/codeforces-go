package main

import (
	. "fmt"
	"io"
	"math"
)

// github.com/EndlessCheng/codeforces-go
func CF453A(in io.Reader, out io.Writer) {
	var m, n int
	Fscan(in, &m, &n)
	ans := float64(m)
	for i := 1; i < m; i++ {
		ans -= math.Pow(float64(i)/float64(m), float64(n))
	}
	Fprintf(out, "%f", ans)
}

//func main() { CF453A(os.Stdin, os.Stdout) }
