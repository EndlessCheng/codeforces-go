package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1351B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var t, a, b, c, d int
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &a, &b, &c, &d)
		if a > b {
			a, b = b, a
		}
		if c > d {
			c, d = d, c
		}
		if b == d && b == a+c {
			Fprintln(out, "Yes")
		} else {
			Fprintln(out, "No")
		}
	}
}

//func main() { CF1351B(os.Stdin, os.Stdout) }
