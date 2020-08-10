package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF101C(in io.Reader, out io.Writer) {
	var x, y [3]int64
	for i := 0; i < 3; i++ {
		Fscan(in, &x[i], &y[i])
	}
	d := x[2]*x[2] + y[2]*y[2]
	for i := 0; i < 4; i++ {
		x[0], y[0] = y[0], -x[0]
		dx, dy := x[1]-x[0], y[1]-y[0]
		d1, d2 := dx*x[2]+dy*y[2], dx*y[2]-dy*x[2]
		if d == 0 && dx == 0 && dy == 0 || d > 0 && d1%d == 0 && d2%d == 0 {
			Fprint(out, "YES")
			return
		}
	}
	Fprint(out, "NO")
}

//func main() { CF101C(os.Stdin, os.Stdout) }
