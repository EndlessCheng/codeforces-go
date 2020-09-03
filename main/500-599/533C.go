package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF533C(in io.Reader, out io.Writer) {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	var x1, y1, x2, y2 int
	Fscan(in, &x1, &y1, &x2, &y2)
	if x1+y1 <= max(x2, y2) || x1 <= x2 && y1 <= y2 {
		Fprint(out, "Polycarp")
	} else {
		Fprint(out, "Vasiliy")
	}
}

//func main() { CF533C(os.Stdin, os.Stdout) }
