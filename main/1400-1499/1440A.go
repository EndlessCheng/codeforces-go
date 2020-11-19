package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1440A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, c0, c1, h int
	var s []byte
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &c0, &c1, &h, &s)
		if c0 > c1+h {
			c0 = c1 + h
		} else if c1 > c0+h {
			c1 = c0 + h
		}
		b := [2]int{}
		for _, c := range s {
			b[c&1]++
		}
		Fprintln(out, c0*b[0]+c1*b[1])
	}
}

//func main() { CF1440A(os.Stdin, os.Stdout) }
