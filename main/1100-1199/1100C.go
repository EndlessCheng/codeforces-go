package main

import (
	. "fmt"
	"io"
	"math"
)

// https://space.bilibili.com/206214
func cf1100C(in io.Reader, out io.Writer) {
	var n, r float64
	Fscan(in, &n, &r)
	Fprintf(out, "%.7f", r/(1/math.Sin(math.Pi/n)-1))
}

//func main() { cf1100C(os.Stdin, os.Stdout) }
