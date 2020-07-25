package main

import (
	. "fmt"
	"io"
	"math"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(in io.Reader, out io.Writer) {
	var n, m int
	Fscan(in, &n, &m)
	if m == 1 {
		Fprint(out, n)
	} else if m == 2 {
		x := int(math.Sqrt(float64(n))) - 1
		for ; x*x <= n; x++ {
		}
		x--
		Fprint(out, x)
	} else if m > 31 {
		Fprint(out, 1)
	} else {
		x := 1
	o:
		for {
			s := 1
			for i := 1; i <= m; i++ {
				s *= x
				if s > n {
					break o
				}
			}
			x++
		}
		Fprint(out, x-1)
	}
}

func main() { run(os.Stdin, os.Stdout) }
