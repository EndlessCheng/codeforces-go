package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var x0, y0, x, y int
	Fscan(in, &x0, &y0, &x, &y)
	x -= x0
	y -= y0
	if x == 0 && y == 0 {
		Fprint(out, 0)
		return
	}
	x, y = abs(x), abs(y)
	if x+y <= 3 || x+y == 0 || x == y {
		Fprint(out, 1)
		return
	}
	if y > x {
		y, x = x, y
	}
	if x-y <= 3 || (x+y)&1 == 0 {
		Fprint(out, 2)
	} else {
		Fprint(out, 3)
	}
}

func main() { run(os.Stdin, os.Stdout) }

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
