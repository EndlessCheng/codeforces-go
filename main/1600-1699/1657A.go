package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
)

// github.com/EndlessCheng/codeforces-go
func CF1657A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, x, y int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &x, &y)
		if x == 0 && y == 0 {
			Fprintln(out, 0)
		} else {
			s := x*x + y*y
			rt := int(math.Sqrt(float64(s)))
			if rt*rt == s {
				Fprintln(out, 1)
			} else {
				Fprintln(out, 2)
			}
		}
	}
}

//func main() { CF1657A(os.Stdin, os.Stdout) }
