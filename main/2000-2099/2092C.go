package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2092C(in io.Reader, out io.Writer) {
	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		var s, mx, c1 int
		for range n {
			Fscan(in, &v)
			s += v
			mx = max(mx, v)
			c1 += v & 1
		}
		if c1 == 0 || c1 == n {
			Fprintln(out, mx)
		} else {
			Fprintln(out, s-c1+1)
		}
	}
}

//func main() { cf2092C(bufio.NewReader(os.Stdin), os.Stdout) }
