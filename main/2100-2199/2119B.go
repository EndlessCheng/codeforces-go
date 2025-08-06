package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2119B(in io.Reader, out io.Writer) {
	var T, n, x, y, tx, ty int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &x, &y, &tx, &ty)
		d := (x-tx)*(x-tx) + (y-ty)*(y-ty)
		s, mx := 0, 0
		for range n {
			Fscan(in, &x)
			s += x
			mx = max(mx, x)
		}
		v := mx*2 - s
		if d > s*s || v > 0 && d < v*v {
			Fprintln(out, "No")
		} else {
			Fprintln(out, "Yes")
		}
	}
}

//func main() { cf2119B(bufio.NewReader(os.Stdin), os.Stdout) }
