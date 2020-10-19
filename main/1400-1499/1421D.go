package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1421D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	min := func(a, b int64) int64 {
		if a < b {
			return a
		}
		return b
	}

	var T, x, y, c1, c2, c3, c4, c5, c6 int64
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &x, &y, &c1, &c2, &c3, &c4, &c5, &c6)
		if x < 0 {
			x, y, c1, c2, c3, c4, c5, c6 = -x, -y, c4, c5, c6, c1, c2, c3
		}
		if y > x {
			c1 = min(c1, c2+c6)
			c2 = min(c2, c1+c3)
			Fprintln(out, x*c1+(y-x)*c2)
		} else if y > 0 {
			x, y, c2, c3, c5, c6 = y, x, c6, c5, c3, c2
			c1 = min(c1, c2+c6)
			c2 = min(c2, c1+c3)
			Fprintln(out, x*c1+(y-x)*c2)
		} else {
			c6 = min(c6, c1+c5)
			c5 = min(c5, c6+c4)
			Fprintln(out, x*c6-y*c5)
		}
	}
}

//func main() { CF1421D(os.Stdin, os.Stdout) }
