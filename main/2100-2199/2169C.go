package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2169C(in io.Reader, out io.Writer) {
	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		s, pre, mx := 0, int(-1e18), 0
		for r := 1; r <= n; r++ {
			Fscan(in, &v)
			pre = max(pre, s+r-r*r)
			s += v
			mx = max(mx, r*r+r-s+pre)
		}
		Fprintln(out, s+mx)
	}
}

//func main() { cf2169C(bufio.NewReader(os.Stdin), os.Stdout) }
