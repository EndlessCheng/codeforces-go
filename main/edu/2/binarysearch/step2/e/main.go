package main

import (
	. "fmt"
	"io"
	"math"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(in io.Reader, out io.Writer) {
	const eps = 1e-8
	var c float64
	Fscan(in, &c)
	l, r := 0.0, math.Sqrt(c)
	for t := int(math.Log2((r - l) / eps)); t > 0; t-- {
		m := (l + r) / 2
		if m*m+math.Sqrt(m) >= c {
			r = m
		} else {
			l = m
		}
	}
	Fprintf(out, "%.8f", (l+r)/2)
}

func main() { run(os.Stdin, os.Stdout) }
