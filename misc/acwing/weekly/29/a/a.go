package main

import (
	. "fmt"
	"io"
	"math"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(in io.Reader, out io.Writer) {
	var n, t float64
	Fscan(in, &n, &t)
	Fprintf(out, "%.6f", n*math.Pow(1.00011, t))
}

func main() { run(os.Stdin, os.Stdout) }
