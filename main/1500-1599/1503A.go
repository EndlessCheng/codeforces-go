package main

import (
	"bufio"
	. "fmt"
	"io"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func CF1503A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	var s string
	var flip bool
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &s)
		c1 := strings.Count(s, "1")
		if s[0] == '0' || s[n-1] == '0' || c1&1 > 0 {
			Fprintln(out, "NO")
			continue
		}
		c1 /= 2
		x := make([]byte, n)
		y := make([]byte, n)
		for i, b := range s {
			if b == '1' {
				if c1 > 0 {
					c1--
					x[i] = '('
					y[i] = '('
				} else {
					x[i] = ')'
					y[i] = ')'
				}
			} else {
				flip = !flip
				if flip {
					x[i] = '('
					y[i] = ')'
				} else {
					x[i] = ')'
					y[i] = '('
				}
			}
		}
		Fprintf(out, "YES\n%s\n%s\n", x, y)
	}
}

//func main() { CF1503A(os.Stdin, os.Stdout) }
