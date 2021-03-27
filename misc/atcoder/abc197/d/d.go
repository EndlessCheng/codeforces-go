package main

import (
	. "fmt"
	"io"
	"math"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(in io.Reader, out io.Writer) {
	var n, x0, y0, xm, ym float64
	Fscan(in, &n, &x0, &y0, &xm, &ym)
	x, y, rad := (x0-xm)/2, (y0-ym)/2, 2*math.Pi/n
	Fprintf(out, "%.11f %.11f", x*math.Cos(rad)-y*math.Sin(rad)+(x0+xm)/2, x*math.Sin(rad)+y*math.Cos(rad)+(y0+ym)/2)
}

func main() { run(os.Stdin, os.Stdout) }
