package main

import (
	"bufio"
	"bytes"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1896F(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n int
	var s []byte
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &s)
		m := n * 2
		if s[0] != s[m-1] || bytes.Count(s, []byte{'1'})%2 > 0 {
			Fprintln(out, -1)
			continue
		}

		if s[0]&1 > 0 {
			Fprintln(out, 3)
			Fprint(out, "(")
			for i := n - 2; i >= 0; i-- {
				Fprint(out, "()")
			}
			Fprintln(out, ")")
			Fprint(out, "(")
			s[0] ^= 1
			s[m-1] ^= 1
		} else {
			Fprintln(out, 2)
			Fprint(out, "(")
		}

		c := 0
		for i := 1; i < m-1; i += 2 {
			if s[i] != s[i+1] {
				c ^= 1
				s[i+c] ^= 1
				if c > 0 {
					Fprint(out, "((")
				} else {
					Fprint(out, "))")
				}
			} else {
				Fprint(out, "()")
			}
		}
		Fprintln(out, ")")

		Fprint(out, "(")
		for i := 1; i < m-1; i += 2 {
			if s[i]&1 > 0 {
				Fprint(out, ")(")
			} else {
				Fprint(out, "()")
			}
		}
		Fprintln(out, ")")
	}
}

//func main() { cf1896F(bufio.NewReader(os.Stdin), os.Stdout) }
