package main

import (
	. "fmt"
	"io"
	"math"
)

// https://space.bilibili.com/206214
func CF407A(in io.Reader, out io.Writer) {
	var a, b int
	Fscan(in, &a, &b)
	for x := 1; x < a; x++ {
		if b*x%a == 0 {
			y := int(math.Sqrt(float64(a*a - x*x)))
			if y*y != a*a-x*x || y == b*x/a {
				continue
			}
			Fprintln(out, "YES")
			Fprintln(out, 0, 0)
			Fprintln(out, x, y)
			Fprintln(out, -b*y/a, b*x/a)
			return
		}
	}
	Fprint(out, "NO")
}

//func main() { CF407A(os.Stdin, os.Stdout) }
