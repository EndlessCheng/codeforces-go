package main

import (
	. "fmt"
	"io"
	"os"
)

// https://github.com/EndlessCheng
func f(x, y int) int {
	if x < 0 || y < 0 {
		return 0
	}
	if y > x {
		y, x = x, y
	}
	k := y/2 + 1
	return k*k*2 - k + (x/2-y/2)*(y+1)
}

func calc(x1, y1, x2, y2 int) int {
	if x2 < 0 {
		x1, x2 = -x2, -x1
	}
	if y2 < 0 {
		y1, y2 = -y2, -y1
	}
	if x1 < 0 {
		return calc(1, y1, -x1, y2) + calc(0, y1, x2, y2)
	}
	if y1 < 0 {
		return calc(x1, 1, x2, -y1) + calc(x1, 0, x2, y2)
	}
	return f(x2, y2) - f(x1-1, y2) - f(x2, y1-1) + f(x1-1, y1-1)
}

func run(in io.Reader, out io.Writer) {
	var l, r, d, u int
	Fscan(in, &l, &r, &d, &u)
	Fprint(out, calc(l, d, r, u))
}

func main() { run(os.Stdin, os.Stdout) }
