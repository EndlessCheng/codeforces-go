package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2096E(in io.Reader, out io.Writer) {
	var T, n int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &s)
		var inv, d, r int
		for _, c := range s {
			if c == 'P' {
				r++
				continue
			}
			inv += r
			if d%2 != r%2 {
				d++
			} else if d > 0 {
				d--
			}
		}
		d = (d + 1) / 2
		Fprintln(out, (inv-d)/2+d)
	}
}

//func main() { cf2096E(bufio.NewReader(os.Stdin), os.Stdout) }
