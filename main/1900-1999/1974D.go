package main

import (
	"bytes"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1974D(in io.Reader, out io.Writer) {
	var T, n int
	var s []byte
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &s)
		w := ['X']int{'W': 1, 'E': 1}
		for i, b := range s {
			s[i] = "RH"[w[b]]
			w[b] ^= 1
		}
		c := bytes.Count(s, []byte{'R'})
		if c == 0 || c == n || w['N'] != w['S'] || w['W'] != w['E'] {
			Fprintln(out, "NO")
		} else {
			Fprintf(out, "%s\n", s)
		}
	}
}

//func main() { cf1974D(bufio.NewReader(os.Stdin), os.Stdout) }
