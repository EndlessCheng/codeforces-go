package main

import (
	. "fmt"
	"io"
	"math"
	"os"
)

// https://github.com/EndlessCheng
func isqrt(x int) int {
	rt := int(math.Sqrt(float64(x)))
	if rt*rt > x {
		rt--
	}
	return rt
}

func run(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	Fprint(out, isqrt(n/2)+isqrt(n/4))
}

func main() { run(os.Stdin, os.Stdout) }
